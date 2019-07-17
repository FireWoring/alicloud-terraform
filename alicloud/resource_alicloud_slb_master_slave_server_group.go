package alicloud

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/slb"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/terraform-providers/terraform-provider-alicloud/alicloud/connectivity"
)

func resourceAliyunSlbMasterSlaveServerGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliyunSlbMasterSlaveServerGroupCreate,
		Read:   resourceAliyunSlbMasterSlaveServerGroupRead,
		//Update: resourceAliyunSlbMasterSlaveServerGroupUpdate,
		Delete: resourceAliyunSlbMasterSlaveServerGroupDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"load_balancer_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				//Default:  "tf-master-slave-server-group",
			},

			"servers": {
				Type:     schema.TypeSet,
				Required: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"server_id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"port": {
							Type:         schema.TypeInt,
							Required:     true,
							ValidateFunc: validateIntegerInRange(1, 65535),
						},
						"weight": {
							Type:         schema.TypeInt,
							Optional:     true,
							Default:      100,
							ValidateFunc: validateIntegerInRange(0, 100),
						},
						"type": {
							Type:         schema.TypeString,
							Optional:     true,
							Default:      string(ECS),
							ValidateFunc: validateAllowedStringValue([]string{string(ENI), string(ECS)}),
						},
						"server_type": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validateAllowedStringValue([]string{string("Master"), string("Slave")}),
						},
						"is_backup": {
							Type:         schema.TypeInt,
							Optional:     true,
							ValidateFunc: validateAllowedIntValue([]int{0, 1}),
						},
					},
				},
				MaxItems: 2,
				MinItems: 2,
			},
		},
	}
}

func resourceAliyunSlbMasterSlaveServerGroupCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	request := slb.CreateCreateMasterSlaveServerGroupRequest()
	request.LoadBalancerId = d.Get("load_balancer_id").(string)
	if v, ok := d.GetOk("name"); ok {
		request.MasterSlaveServerGroupName = v.(string)
	}
	if v, ok := d.GetOk("servers"); ok {
		request.MasterSlaveBackendServers = expandMasterSlaveBackendServersToString(v.(*schema.Set).List())
	}
	raw, err := client.WithSlbClient(func(slbClient *slb.Client) (interface{}, error) {
		return slbClient.CreateMasterSlaveServerGroup(request)
	})
	if err != nil {
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_slb_master_slave_server_group", request.GetActionName(), AlibabaCloudSdkGoERROR)
	}
	addDebug(request.GetActionName(), raw)
	response, _ := raw.(*slb.CreateMasterSlaveServerGroupResponse)
	d.SetId(response.MasterSlaveServerGroupId)

	return resourceAliyunSlbMasterSlaveServerGroupRead(d, meta)
}

func resourceAliyunSlbMasterSlaveServerGroupRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	slbService := SlbService{client}
	object, err := slbService.DescribeSlbMasterSlaveServerGroup(d.Id())

	if err != nil {
		if NotFoundError(err) {
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("name", object.MasterSlaveServerGroupName)
	//d.Set("load_balancer_id", object.)

	servers := make([]map[string]interface{}, 0)
	portAndWeight := make(map[string]string)
	for _, server := range object.MasterSlaveBackendServers.MasterSlaveBackendServer {
		key := fmt.Sprintf("%d%s%d%s%s%s%s%s%d", server.Port, COLON_SEPARATED, server.Weight, COLON_SEPARATED, server.Type, COLON_SEPARATED, server.ServerType, COLON_SEPARATED, server.IsBackup)

		portAndWeight[key] = server.ServerId
	}
	for key, value := range portAndWeight {
		k := strings.Split(key, COLON_SEPARATED)
		p, e := strconv.Atoi(k[0])
		if e != nil {
			return WrapError(e)
		}
		w, e := strconv.Atoi(k[1])
		if e != nil {
			return WrapError(e)
		}
		t := k[2]
		st := k[3]
		isBackup, e := strconv.Atoi(k[4])
		if e != nil {
			return WrapError(e)
		}
		s := map[string]interface{}{
			"server_id":   value,
			"port":        p,
			"weight":      w,
			"type":        t,
			"server_type": st,
			"is_backup":   isBackup,
		}
		servers = append(servers, s)
	}

	if err := d.Set("servers", servers); err != nil {
		return WrapError(err)
	}

	return nil
}

func resourceAliyunSlbMasterSlaveServerGroupDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	slbService := SlbService{client}
	request := slb.CreateDeleteMasterSlaveServerGroupRequest()
	request.MasterSlaveServerGroupId = d.Id()

	err := resource.Retry(5*time.Minute, func() *resource.RetryError {
		raw, err := client.WithSlbClient(func(slbClient *slb.Client) (interface{}, error) {
			return slbClient.DeleteMasterSlaveServerGroup(request)
		})
		if err != nil {
			if IsExceptedErrors(err, []string{RspoolVipExist}) {
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		addDebug(request.GetActionName(), raw)
		return nil
	})
	if err != nil {
		if IsExceptedErrors(err, []string{MasterSlaveServerGroupNotFoundMessage, InvalidParameter}) {
			return nil
		}
		return WrapErrorf(err, DefaultErrorMsg, d.Id(), request.GetActionName(), AlibabaCloudSdkGoERROR)
	}

	return WrapError(slbService.WaitForSlbMasterSlaveServerGroup(d.Id(), Deleted, DefaultTimeoutMedium))
}

package metrics

import (
	"encoding/json"
	"log"
)

// HandleCollectionRequest takes the collection request payload and updates the target credential info if needed
func HandleCollectionRequest(in []byte) []byte {
	ireq := &CollectionRequest{}

	e := json.Unmarshal(in, ireq)
	if nil != e {
		log.Println("Error in unmarshal: " + e.Error())
		return in
	}

	// Query the Credential manager only if we get a reference to access_key
	if "access_key_ref" != ireq.CollectorPayload.Target.Credentials.Type {
		return in
	}

	details := &ireq.CollectorPayload.Target.Credentials.Details

	accessKeyRefVal := ""
	for _, v := range *details {
		if v.Name == "access_key_id" {
			accessKeyRefVal = v.Value
			log.Println("Access key ref is:", accessKeyRefVal)
		}
	}
	// map[domain:www.microfocus.com filterConfig:map[resourceList:[arn:aws:ec2:us-east-1:040643014701:instance/i-01381e5215e1be55f] serviceList:[ec2]
	// targetTags:[map[name:App value:CollaborationPortal]]] jobId:testJobID jobName:Cyclos Banking on AWS metricConfig:map[expression:[ec2/CPUUtilization]
	// metricList:[]] periodSecs:600 scopeConfig:map[entityList:[] entityQuery:] startTimeOffsetInSeconds:600
	// target:map[count:3 credentials:map[details:[map[name:access_key_id value:test_access_key]]
	// id:a250eaa8-2c29-4d7d-a46e-7fc33bf1431d name:Cyclos Prod Monitoring Account type:access_key_ref]
	// description:1 endpoint:www.microfocus.com host:us-east-1 id:8eff03c1-dad9-48f4-9d9e-a5cfd286f3ac
	// name:AWS us-east-1 protocol:http proxyCredentials:map[credential: id: name: principal: type:]
	// proxyUrl:http://web-proxy.in.softwaregrp.net:8080/ requestTimeoutInSeconds:300 type:aws]]

	// Will be something like:  map[name:access_key_id value:test_access_key]
	// Need to get the value for access key ref
	// Check if name = access_key_id in the previous for loop, if yes, get the value
	newDetails := []CredentialDetailConfig{}

	// TODO: Resolve and append the name value pairs returned by Credential manager
	e1 := CredentialDetailConfig{"access_key_id", "id1"}
	newDetails = append(newDetails, e1)
	e1 = CredentialDetailConfig{"secret_access_key", "sk1"}
	newDetails = append(newDetails, e1)

	log.Println("original json: ", string(in))
	// Update from access key ref to actual access key

	*details = append(*details, newDetails...)
	nb, e := json.Marshal(&ireq)
	log.Println("Updated json: ", string(nb))
	return nil
}

// Sample updated payload with access_key
// map[domain:www.microfocus.com filterConfig:map[resourceList:[arn:aws:ec2:us-east-1:040643014701:instance/i-01381e5215e1be55f] serviceList:[ec2] targetTags:[map[name:App value:CollaborationPortal]]] jobId:testJobID jobName:Cyclos Banking on AWS metricConfig:map[expression:[ec2/CPUUtilization] metricList:[]] periodSecs:600 scopeConfig:map[entityList:[] entityQuery:] startTimeOffsetInSeconds:600 target:map[count:3 credentials:map[details:[{access_key_id id1} {secret_access_key sk1}] id:a250eaa8-2c29-4d7d-a46e-7fc33bf1431d name:Cyclos Prod Monitoring Account type:access_key] description:1 endpoint:www.microfocus.com host:us-east-1 id:8eff03c1-dad9-48f4-9d9e-a5cfd286f3ac name:AWS us-east-1 protocol:http proxyCredentials:map[credential: id: name: principal: type:] proxyUrl:http://web-proxy.in.softwaregrp.net:8080/ requestTimeoutInSeconds:300 type:aws]]

func getObject(name string, fromObj map[string]interface{}) map[string]interface{} {
	var obj map[string]interface{}
	for k, v := range fromObj {
		if k == name {
			obj = v.(map[string]interface{})
			break
		}
	}
	return obj
}

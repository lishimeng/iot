package nwconnector

const (
	topicSubTpl = "schedule/network/%d/+/+"   // schedule/network/{network_type}/{network_id}/{event_type}
	topicPubTpl = "schedule/network/%d/%s/%d" // schedule/network/{network_type}/{network_id}/{event_type}
)

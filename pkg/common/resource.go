package shared

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

type PlayerIDType uuid.UUID

const (
	ResourceTypeGroup          ResourceType = "Groups"         // system group
	ResourceTypePlan           ResourceType = "Plans"          //
	ResourceTypeSubscription   ResourceType = "Subscriptions"  //
	ResourceTypeUser           ResourceType = "Users"          // specification of group
	ResourceTypePage           ResourceType = "Pages"          //
	ResourceTypeList           ResourceType = "List"           // recurse root resources (?)
	ResourceTypePayment        ResourceType = "Payments"
	ResourceTypeWallet         ResourceType = "Wallets"
	ResourceTypeTag            ResourceType = "Tags"
)

var ResourceTypes = []ResourceType{
	ResourceTypeGroup,
	ResourceTypePlan,
	ResourceTypeSubscription,
	ResourceTypeUser,
	ResourceTypePage,
	ResourceTypeList,
	ResourceTypePayment,
	ResourceTypeWallet,
	ResourceTypeTag,
}

var ResourceKeyMap = map[ResourceType]string{
	ResourceTypeGroup: "group_id",
	ResourceTypeUser:  "user_id",
}

func GetResourceFieldID(resourcePart string) (string, error) {
	for k, v := range ResourceKeyMap {
		if strings.EqualFold(fmt.Sprint(k), resourcePart) {
			return v, nil
		}
	}

	return "", fmt.Errorf("failed to parse ResourceIDField: Unknown resource %s", resourcePart)
}

type Resource struct {
	ID   uuid.UUID    `json:"id" bson:"_id"`
	Type ResourceType `json:"type" bson:"type"`
	// ResourceSlug string       `json:"resource_slug" bson:"resource_slug"`
}

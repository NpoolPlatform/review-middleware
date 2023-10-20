// Code generated by ent, DO NOT EDIT.

//go:build tools
// +build tools

// Package internal holds a loadable version of the latest schema.
package internal

const Schema = `{"Schema":"github.com/NpoolPlatform/review-middleware/pkg/db/ent/schema","Package":"github.com/NpoolPlatform/review-middleware/pkg/db/ent","Schemas":[{"name":"PubsubMessage","config":{"Table":""},"fields":[{"name":"created_at","type":{"Type":16,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"default":true,"default_kind":19,"position":{"Index":0,"MixedIn":true,"MixinIndex":0}},{"name":"updated_at","type":{"Type":16,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"default":true,"default_kind":19,"update_default":true,"position":{"Index":1,"MixedIn":true,"MixinIndex":0}},{"name":"deleted_at","type":{"Type":16,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"default":true,"default_kind":19,"position":{"Index":2,"MixedIn":true,"MixinIndex":0}},{"name":"id","type":{"Type":4,"Ident":"uuid.UUID","PkgPath":"github.com/google/uuid","PkgName":"uuid","Nillable":false,"RType":{"Name":"UUID","Ident":"uuid.UUID","Kind":17,"PkgPath":"github.com/google/uuid","Methods":{"ClockSequence":{"In":[],"Out":[{"Name":"int","Ident":"int","Kind":2,"PkgPath":"","Methods":null}]},"Domain":{"In":[],"Out":[{"Name":"Domain","Ident":"uuid.Domain","Kind":8,"PkgPath":"github.com/google/uuid","Methods":null}]},"ID":{"In":[],"Out":[{"Name":"uint32","Ident":"uint32","Kind":10,"PkgPath":"","Methods":null}]},"MarshalBinary":{"In":[],"Out":[{"Name":"","Ident":"[]uint8","Kind":23,"PkgPath":"","Methods":null},{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"MarshalText":{"In":[],"Out":[{"Name":"","Ident":"[]uint8","Kind":23,"PkgPath":"","Methods":null},{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"NodeID":{"In":[],"Out":[{"Name":"","Ident":"[]uint8","Kind":23,"PkgPath":"","Methods":null}]},"Scan":{"In":[{"Name":"","Ident":"interface {}","Kind":20,"PkgPath":"","Methods":null}],"Out":[{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"String":{"In":[],"Out":[{"Name":"string","Ident":"string","Kind":24,"PkgPath":"","Methods":null}]},"Time":{"In":[],"Out":[{"Name":"Time","Ident":"uuid.Time","Kind":6,"PkgPath":"github.com/google/uuid","Methods":null}]},"URN":{"In":[],"Out":[{"Name":"string","Ident":"string","Kind":24,"PkgPath":"","Methods":null}]},"UnmarshalBinary":{"In":[{"Name":"","Ident":"[]uint8","Kind":23,"PkgPath":"","Methods":null}],"Out":[{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"UnmarshalText":{"In":[{"Name":"","Ident":"[]uint8","Kind":23,"PkgPath":"","Methods":null}],"Out":[{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"Value":{"In":[],"Out":[{"Name":"Value","Ident":"driver.Value","Kind":20,"PkgPath":"database/sql/driver","Methods":null},{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"Variant":{"In":[],"Out":[{"Name":"Variant","Ident":"uuid.Variant","Kind":8,"PkgPath":"github.com/google/uuid","Methods":null}]},"Version":{"In":[],"Out":[{"Name":"Version","Ident":"uuid.Version","Kind":8,"PkgPath":"github.com/google/uuid","Methods":null}]}}}},"unique":true,"position":{"Index":0,"MixedIn":false,"MixinIndex":0}},{"name":"message_id","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"default":true,"default_value":"DefaultMsgID","default_kind":24,"position":{"Index":1,"MixedIn":false,"MixinIndex":0}},{"name":"state","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"default":true,"default_value":"DefaultMsgState","default_kind":24,"position":{"Index":2,"MixedIn":false,"MixinIndex":0}},{"name":"resp_to_id","type":{"Type":4,"Ident":"uuid.UUID","PkgPath":"github.com/google/uuid","PkgName":"uuid","Nillable":false,"RType":{"Name":"UUID","Ident":"uuid.UUID","Kind":17,"PkgPath":"github.com/google/uuid","Methods":{"ClockSequence":{"In":[],"Out":[{"Name":"int","Ident":"int","Kind":2,"PkgPath":"","Methods":null}]},"Domain":{"In":[],"Out":[{"Name":"Domain","Ident":"uuid.Domain","Kind":8,"PkgPath":"github.com/google/uuid","Methods":null}]},"ID":{"In":[],"Out":[{"Name":"uint32","Ident":"uint32","Kind":10,"PkgPath":"","Methods":null}]},"MarshalBinary":{"In":[],"Out":[{"Name":"","Ident":"[]uint8","Kind":23,"PkgPath":"","Methods":null},{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"MarshalText":{"In":[],"Out":[{"Name":"","Ident":"[]uint8","Kind":23,"PkgPath":"","Methods":null},{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"NodeID":{"In":[],"Out":[{"Name":"","Ident":"[]uint8","Kind":23,"PkgPath":"","Methods":null}]},"Scan":{"In":[{"Name":"","Ident":"interface {}","Kind":20,"PkgPath":"","Methods":null}],"Out":[{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"String":{"In":[],"Out":[{"Name":"string","Ident":"string","Kind":24,"PkgPath":"","Methods":null}]},"Time":{"In":[],"Out":[{"Name":"Time","Ident":"uuid.Time","Kind":6,"PkgPath":"github.com/google/uuid","Methods":null}]},"URN":{"In":[],"Out":[{"Name":"string","Ident":"string","Kind":24,"PkgPath":"","Methods":null}]},"UnmarshalBinary":{"In":[{"Name":"","Ident":"[]uint8","Kind":23,"PkgPath":"","Methods":null}],"Out":[{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"UnmarshalText":{"In":[{"Name":"","Ident":"[]uint8","Kind":23,"PkgPath":"","Methods":null}],"Out":[{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"Value":{"In":[],"Out":[{"Name":"Value","Ident":"driver.Value","Kind":20,"PkgPath":"database/sql/driver","Methods":null},{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"Variant":{"In":[],"Out":[{"Name":"Variant","Ident":"uuid.Variant","Kind":8,"PkgPath":"github.com/google/uuid","Methods":null}]},"Version":{"In":[],"Out":[{"Name":"Version","Ident":"uuid.Version","Kind":8,"PkgPath":"github.com/google/uuid","Methods":null}]}}}},"optional":true,"default":true,"default_kind":19,"position":{"Index":3,"MixedIn":false,"MixinIndex":0}},{"name":"undo_id","type":{"Type":4,"Ident":"uuid.UUID","PkgPath":"github.com/google/uuid","PkgName":"uuid","Nillable":false,"RType":{"Name":"UUID","Ident":"uuid.UUID","Kind":17,"PkgPath":"github.com/google/uuid","Methods":{"ClockSequence":{"In":[],"Out":[{"Name":"int","Ident":"int","Kind":2,"PkgPath":"","Methods":null}]},"Domain":{"In":[],"Out":[{"Name":"Domain","Ident":"uuid.Domain","Kind":8,"PkgPath":"github.com/google/uuid","Methods":null}]},"ID":{"In":[],"Out":[{"Name":"uint32","Ident":"uint32","Kind":10,"PkgPath":"","Methods":null}]},"MarshalBinary":{"In":[],"Out":[{"Name":"","Ident":"[]uint8","Kind":23,"PkgPath":"","Methods":null},{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"MarshalText":{"In":[],"Out":[{"Name":"","Ident":"[]uint8","Kind":23,"PkgPath":"","Methods":null},{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"NodeID":{"In":[],"Out":[{"Name":"","Ident":"[]uint8","Kind":23,"PkgPath":"","Methods":null}]},"Scan":{"In":[{"Name":"","Ident":"interface {}","Kind":20,"PkgPath":"","Methods":null}],"Out":[{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"String":{"In":[],"Out":[{"Name":"string","Ident":"string","Kind":24,"PkgPath":"","Methods":null}]},"Time":{"In":[],"Out":[{"Name":"Time","Ident":"uuid.Time","Kind":6,"PkgPath":"github.com/google/uuid","Methods":null}]},"URN":{"In":[],"Out":[{"Name":"string","Ident":"string","Kind":24,"PkgPath":"","Methods":null}]},"UnmarshalBinary":{"In":[{"Name":"","Ident":"[]uint8","Kind":23,"PkgPath":"","Methods":null}],"Out":[{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"UnmarshalText":{"In":[{"Name":"","Ident":"[]uint8","Kind":23,"PkgPath":"","Methods":null}],"Out":[{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"Value":{"In":[],"Out":[{"Name":"Value","Ident":"driver.Value","Kind":20,"PkgPath":"database/sql/driver","Methods":null},{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"Variant":{"In":[],"Out":[{"Name":"Variant","Ident":"uuid.Variant","Kind":8,"PkgPath":"github.com/google/uuid","Methods":null}]},"Version":{"In":[],"Out":[{"Name":"Version","Ident":"uuid.Version","Kind":8,"PkgPath":"github.com/google/uuid","Methods":null}]}}}},"optional":true,"default":true,"default_kind":19,"position":{"Index":4,"MixedIn":false,"MixinIndex":0}},{"name":"arguments","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"size":2147483647,"optional":true,"default":true,"default_value":"","default_kind":24,"position":{"Index":5,"MixedIn":false,"MixinIndex":0}}],"indexes":[{"fields":["state","resp_to_id"]},{"fields":["state","undo_id"]}],"policy":[{"Index":0,"MixedIn":true,"MixinIndex":0}]},{"name":"Review","config":{"Table":""},"fields":[{"name":"created_at","type":{"Type":16,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"default":true,"default_kind":19,"position":{"Index":0,"MixedIn":true,"MixinIndex":0}},{"name":"updated_at","type":{"Type":16,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"default":true,"default_kind":19,"update_default":true,"position":{"Index":1,"MixedIn":true,"MixinIndex":0}},{"name":"deleted_at","type":{"Type":16,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"default":true,"default_kind":19,"position":{"Index":2,"MixedIn":true,"MixinIndex":0}},{"name":"id","type":{"Type":16,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":0,"MixedIn":true,"MixinIndex":1}},{"name":"ent_id","type":{"Type":4,"Ident":"uuid.UUID","PkgPath":"github.com/google/uuid","PkgName":"uuid","Nillable":false,"RType":{"Name":"UUID","Ident":"uuid.UUID","Kind":17,"PkgPath":"github.com/google/uuid","Methods":{"ClockSequence":{"In":[],"Out":[{"Name":"int","Ident":"int","Kind":2,"PkgPath":"","Methods":null}]},"Domain":{"In":[],"Out":[{"Name":"Domain","Ident":"uuid.Domain","Kind":8,"PkgPath":"github.com/google/uuid","Methods":null}]},"ID":{"In":[],"Out":[{"Name":"uint32","Ident":"uint32","Kind":10,"PkgPath":"","Methods":null}]},"MarshalBinary":{"In":[],"Out":[{"Name":"","Ident":"[]uint8","Kind":23,"PkgPath":"","Methods":null},{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"MarshalText":{"In":[],"Out":[{"Name":"","Ident":"[]uint8","Kind":23,"PkgPath":"","Methods":null},{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"NodeID":{"In":[],"Out":[{"Name":"","Ident":"[]uint8","Kind":23,"PkgPath":"","Methods":null}]},"Scan":{"In":[{"Name":"","Ident":"interface {}","Kind":20,"PkgPath":"","Methods":null}],"Out":[{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"String":{"In":[],"Out":[{"Name":"string","Ident":"string","Kind":24,"PkgPath":"","Methods":null}]},"Time":{"In":[],"Out":[{"Name":"Time","Ident":"uuid.Time","Kind":6,"PkgPath":"github.com/google/uuid","Methods":null}]},"URN":{"In":[],"Out":[{"Name":"string","Ident":"string","Kind":24,"PkgPath":"","Methods":null}]},"UnmarshalBinary":{"In":[{"Name":"","Ident":"[]uint8","Kind":23,"PkgPath":"","Methods":null}],"Out":[{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"UnmarshalText":{"In":[{"Name":"","Ident":"[]uint8","Kind":23,"PkgPath":"","Methods":null}],"Out":[{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"Value":{"In":[],"Out":[{"Name":"Value","Ident":"driver.Value","Kind":20,"PkgPath":"database/sql/driver","Methods":null},{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"Variant":{"In":[],"Out":[{"Name":"Variant","Ident":"uuid.Variant","Kind":8,"PkgPath":"github.com/google/uuid","Methods":null}]},"Version":{"In":[],"Out":[{"Name":"Version","Ident":"uuid.Version","Kind":8,"PkgPath":"github.com/google/uuid","Methods":null}]}}}},"unique":true,"default":true,"default_kind":19,"position":{"Index":1,"MixedIn":true,"MixinIndex":1}},{"name":"app_id","type":{"Type":4,"Ident":"uuid.UUID","PkgPath":"github.com/google/uuid","PkgName":"uuid","Nillable":false,"RType":{"Name":"UUID","Ident":"uuid.UUID","Kind":17,"PkgPath":"github.com/google/uuid","Methods":{"ClockSequence":{"In":[],"Out":[{"Name":"int","Ident":"int","Kind":2,"PkgPath":"","Methods":null}]},"Domain":{"In":[],"Out":[{"Name":"Domain","Ident":"uuid.Domain","Kind":8,"PkgPath":"github.com/google/uuid","Methods":null}]},"ID":{"In":[],"Out":[{"Name":"uint32","Ident":"uint32","Kind":10,"PkgPath":"","Methods":null}]},"MarshalBinary":{"In":[],"Out":[{"Name":"","Ident":"[]uint8","Kind":23,"PkgPath":"","Methods":null},{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"MarshalText":{"In":[],"Out":[{"Name":"","Ident":"[]uint8","Kind":23,"PkgPath":"","Methods":null},{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"NodeID":{"In":[],"Out":[{"Name":"","Ident":"[]uint8","Kind":23,"PkgPath":"","Methods":null}]},"Scan":{"In":[{"Name":"","Ident":"interface {}","Kind":20,"PkgPath":"","Methods":null}],"Out":[{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"String":{"In":[],"Out":[{"Name":"string","Ident":"string","Kind":24,"PkgPath":"","Methods":null}]},"Time":{"In":[],"Out":[{"Name":"Time","Ident":"uuid.Time","Kind":6,"PkgPath":"github.com/google/uuid","Methods":null}]},"URN":{"In":[],"Out":[{"Name":"string","Ident":"string","Kind":24,"PkgPath":"","Methods":null}]},"UnmarshalBinary":{"In":[{"Name":"","Ident":"[]uint8","Kind":23,"PkgPath":"","Methods":null}],"Out":[{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"UnmarshalText":{"In":[{"Name":"","Ident":"[]uint8","Kind":23,"PkgPath":"","Methods":null}],"Out":[{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"Value":{"In":[],"Out":[{"Name":"Value","Ident":"driver.Value","Kind":20,"PkgPath":"database/sql/driver","Methods":null},{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"Variant":{"In":[],"Out":[{"Name":"Variant","Ident":"uuid.Variant","Kind":8,"PkgPath":"github.com/google/uuid","Methods":null}]},"Version":{"In":[],"Out":[{"Name":"Version","Ident":"uuid.Version","Kind":8,"PkgPath":"github.com/google/uuid","Methods":null}]}}}},"optional":true,"default":true,"default_kind":19,"position":{"Index":0,"MixedIn":false,"MixinIndex":0}},{"name":"reviewer_id","type":{"Type":4,"Ident":"uuid.UUID","PkgPath":"github.com/google/uuid","PkgName":"uuid","Nillable":false,"RType":{"Name":"UUID","Ident":"uuid.UUID","Kind":17,"PkgPath":"github.com/google/uuid","Methods":{"ClockSequence":{"In":[],"Out":[{"Name":"int","Ident":"int","Kind":2,"PkgPath":"","Methods":null}]},"Domain":{"In":[],"Out":[{"Name":"Domain","Ident":"uuid.Domain","Kind":8,"PkgPath":"github.com/google/uuid","Methods":null}]},"ID":{"In":[],"Out":[{"Name":"uint32","Ident":"uint32","Kind":10,"PkgPath":"","Methods":null}]},"MarshalBinary":{"In":[],"Out":[{"Name":"","Ident":"[]uint8","Kind":23,"PkgPath":"","Methods":null},{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"MarshalText":{"In":[],"Out":[{"Name":"","Ident":"[]uint8","Kind":23,"PkgPath":"","Methods":null},{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"NodeID":{"In":[],"Out":[{"Name":"","Ident":"[]uint8","Kind":23,"PkgPath":"","Methods":null}]},"Scan":{"In":[{"Name":"","Ident":"interface {}","Kind":20,"PkgPath":"","Methods":null}],"Out":[{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"String":{"In":[],"Out":[{"Name":"string","Ident":"string","Kind":24,"PkgPath":"","Methods":null}]},"Time":{"In":[],"Out":[{"Name":"Time","Ident":"uuid.Time","Kind":6,"PkgPath":"github.com/google/uuid","Methods":null}]},"URN":{"In":[],"Out":[{"Name":"string","Ident":"string","Kind":24,"PkgPath":"","Methods":null}]},"UnmarshalBinary":{"In":[{"Name":"","Ident":"[]uint8","Kind":23,"PkgPath":"","Methods":null}],"Out":[{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"UnmarshalText":{"In":[{"Name":"","Ident":"[]uint8","Kind":23,"PkgPath":"","Methods":null}],"Out":[{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"Value":{"In":[],"Out":[{"Name":"Value","Ident":"driver.Value","Kind":20,"PkgPath":"database/sql/driver","Methods":null},{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"Variant":{"In":[],"Out":[{"Name":"Variant","Ident":"uuid.Variant","Kind":8,"PkgPath":"github.com/google/uuid","Methods":null}]},"Version":{"In":[],"Out":[{"Name":"Version","Ident":"uuid.Version","Kind":8,"PkgPath":"github.com/google/uuid","Methods":null}]}}}},"optional":true,"default":true,"default_kind":19,"position":{"Index":1,"MixedIn":false,"MixinIndex":0}},{"name":"domain","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"default":true,"default_value":"","default_kind":24,"position":{"Index":2,"MixedIn":false,"MixinIndex":0}},{"name":"object_id","type":{"Type":4,"Ident":"uuid.UUID","PkgPath":"github.com/google/uuid","PkgName":"uuid","Nillable":false,"RType":{"Name":"UUID","Ident":"uuid.UUID","Kind":17,"PkgPath":"github.com/google/uuid","Methods":{"ClockSequence":{"In":[],"Out":[{"Name":"int","Ident":"int","Kind":2,"PkgPath":"","Methods":null}]},"Domain":{"In":[],"Out":[{"Name":"Domain","Ident":"uuid.Domain","Kind":8,"PkgPath":"github.com/google/uuid","Methods":null}]},"ID":{"In":[],"Out":[{"Name":"uint32","Ident":"uint32","Kind":10,"PkgPath":"","Methods":null}]},"MarshalBinary":{"In":[],"Out":[{"Name":"","Ident":"[]uint8","Kind":23,"PkgPath":"","Methods":null},{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"MarshalText":{"In":[],"Out":[{"Name":"","Ident":"[]uint8","Kind":23,"PkgPath":"","Methods":null},{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"NodeID":{"In":[],"Out":[{"Name":"","Ident":"[]uint8","Kind":23,"PkgPath":"","Methods":null}]},"Scan":{"In":[{"Name":"","Ident":"interface {}","Kind":20,"PkgPath":"","Methods":null}],"Out":[{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"String":{"In":[],"Out":[{"Name":"string","Ident":"string","Kind":24,"PkgPath":"","Methods":null}]},"Time":{"In":[],"Out":[{"Name":"Time","Ident":"uuid.Time","Kind":6,"PkgPath":"github.com/google/uuid","Methods":null}]},"URN":{"In":[],"Out":[{"Name":"string","Ident":"string","Kind":24,"PkgPath":"","Methods":null}]},"UnmarshalBinary":{"In":[{"Name":"","Ident":"[]uint8","Kind":23,"PkgPath":"","Methods":null}],"Out":[{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"UnmarshalText":{"In":[{"Name":"","Ident":"[]uint8","Kind":23,"PkgPath":"","Methods":null}],"Out":[{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"Value":{"In":[],"Out":[{"Name":"Value","Ident":"driver.Value","Kind":20,"PkgPath":"database/sql/driver","Methods":null},{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"Variant":{"In":[],"Out":[{"Name":"Variant","Ident":"uuid.Variant","Kind":8,"PkgPath":"github.com/google/uuid","Methods":null}]},"Version":{"In":[],"Out":[{"Name":"Version","Ident":"uuid.Version","Kind":8,"PkgPath":"github.com/google/uuid","Methods":null}]}}}},"optional":true,"default":true,"default_kind":19,"position":{"Index":3,"MixedIn":false,"MixinIndex":0}},{"name":"trigger","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"default":true,"default_value":"DefaultTriggerType","default_kind":24,"position":{"Index":4,"MixedIn":false,"MixinIndex":0}},{"name":"object_type","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"default":true,"default_value":"DefaultObjectType","default_kind":24,"position":{"Index":5,"MixedIn":false,"MixinIndex":0}},{"name":"state","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"default":true,"default_value":"DefaultReviewState","default_kind":24,"position":{"Index":6,"MixedIn":false,"MixinIndex":0}},{"name":"message","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"default":true,"default_value":"","default_kind":24,"position":{"Index":7,"MixedIn":false,"MixinIndex":0}}],"indexes":[{"unique":true,"fields":["ent_id"]}],"policy":[{"Index":0,"MixedIn":true,"MixinIndex":0}]}],"Features":["entql","sql/lock","sql/execquery","sql/upsert","privacy","schema/snapshot","sql/modifier"]}`

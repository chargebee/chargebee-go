package enum

type HierarchyOperationType string

const (
	HierarchyOperationTypeCompleteHierarchy HierarchyOperationType = "complete_hierarchy"
	HierarchyOperationTypeSubordinates      HierarchyOperationType = "subordinates"
	HierarchyOperationTypePathToRoot        HierarchyOperationType = "path_to_root"
)

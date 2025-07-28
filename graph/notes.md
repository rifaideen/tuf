# Graph

Types:

1. un-directed graph (edges doesn't have direction specified, can move between both direction between two nodes)
2. directed graph (edges have direction specified, can move in the specified direction only)

**vertices:** node with value

**edges:** connection between nodes

**undirected edge:** edge does not specify the direction

**directed edge:** edge that specify the direction

**cyclic graph:** Graph with cycle, where I can traverse and reach the same node where I started

**acyclic graph:** Graph without cycle, I cannot reach back the same node after traversal

**DAG:** Directed Acyclic Graph

**Path:** contains lot of vertices (nodes) and each of them are reachable. *path cannot have node appearing twice, which is invalid. It can appear only once in a path


**Degrees of graph: Un-Directed Graph**

* number of edges attached to the node.
* total degree of a graph is equal to twice of the number of edges. i.e 2*E

**Degrees of graph: Directed Graph**

* in-degree of node, the number of incoming edges to the node
* out-degree of node, the number of outgoing edges from the ndoe

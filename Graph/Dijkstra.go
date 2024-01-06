package Graph 
import "github.com/paerarason/go-stack/heap"

type edges struct {
	node string 
	distance int8 
} 

type Graphs struct {
	nodes map[string][] edges
} 

func (g *Graphs) AddEdges(start,end string ,dis int8 ){
	 g.nodes[start]=append(g.nodes[start],edges{node:end,distance:dis})
	 g.nodes[end]=append(g.nodes[end],edges{node:start,distance:dis})
}

func DijkstraPath() int,string {
    
}
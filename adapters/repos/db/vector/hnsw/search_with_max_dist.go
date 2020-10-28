package hnsw

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/semi-technologies/weaviate/adapters/repos/db/inverted"
)

func (h *hnsw) knnSearchByVectorMaxDist(searchVec []float32, dist float32,
	ef int, allowList inverted.AllowList) ([]int, error) {
	entryPointID := h.entryPointID
	entryPointDistance, ok, err := h.distBetweenNodeAndVec(entryPointID, searchVec)
	if err != nil {
		return nil, errors.Wrap(err, "knn search: distance between entrypint and query node")
	}

	if !ok {
		return nil, fmt.Errorf("entrypoint was deleted in the object strore, " +
			"it has been flagged for cleanup and should be fixed in the next cleanup cycle")
	}

	// stop at layer 1, not 0!
	for level := h.currentMaximumLayer; level >= 1; level-- {
		eps := &binarySearchTreeGeneric{}
		eps.insert(entryPointID, entryPointDistance)
		// ignore allowList on layers > 0
		res, err := h.searchLayerByVector(searchVec, *eps, 1, level, nil)
		if err != nil {
			return nil, errors.Wrapf(err, "knn search: search layer at level %d", level)
		}
		best := res.minimum()
		entryPointID = best.index
		entryPointDistance = best.dist
	}

	eps := &binarySearchTreeGeneric{}
	eps.insert(entryPointID, entryPointDistance)
	res, err := h.searchLayerByVector(searchVec, *eps, ef, 0, allowList)
	if err != nil {
		return nil, errors.Wrapf(err, "knn search: search layer at level %d", 0)
	}

	flat := res.flattenInOrder()
	out := make([]int, len(flat))
	i := 0
	for _, elem := range flat {
		if elem.dist > dist {
			break
		}
		out[i] = elem.index
		i++
	}

	return out[:i], nil
}

package hierarchy

import (
	"sort"
	"time"

	"github.com/AlexL70/multithreadingingo/deadlocks_train/common"
)

func lockIntersectionInDistance(id, reserveStart, reserveEnd int, crossings []*common.Crossing) {
	var intersectionsToLock []*common.Intersection
	for _, crossing := range crossings {
		if reserveEnd >= crossing.Position && reserveStart <= crossing.Position && crossing.Intersection.LockedBy != id {
			intersectionsToLock = append(intersectionsToLock, crossing.Intersection)
		}
	}
	sort.Slice(intersectionsToLock, func(i, j int) bool {
		return intersectionsToLock[i].Id < intersectionsToLock[j].Id
	})
	for _, it := range intersectionsToLock {
		it.Mutex.Lock()
		it.LockedBy = id
		time.Sleep(10 * time.Millisecond)
	}
}

func MoveTrain(train *common.Train, distance int, crossings []*common.Crossing) {
	for train.Front < distance {
		train.Front++
		for _, crossing := range crossings {
			if train.Front == crossing.Position {
				lockIntersectionInDistance(train.Id, crossing.Position, crossing.Position+train.TrainLength, crossings)
			}
			back := train.Front - train.TrainLength
			if back == crossing.Position {
				crossing.Intersection.LockedBy = -1
				crossing.Intersection.Mutex.Unlock()
			}
		}
		time.Sleep(30 * time.Millisecond)
	}
}

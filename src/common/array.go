package common

func AreIntArrayEqual(array1 []int, array2 []int) bool {
	if len(array1) != len(array2) {
		return false
	} else {
		for index := 0; index <= len(array1); index++ {
			if array1[index] == array2[index] {
				continue
			} else {
				return false
			}
		}
		return true
	}
}

func IsEqual(array1 []Generic, array2 []Generic) bool {
	if len(array1) != len(array2) {
		return false
	} else {
		for index := 0; index <= len(array1); index++ {
			if array1[index] == array2[index] {
				continue
			} else {
				return false
			}
		}
		return true
	}
}

type Generic interface{}
type GenericMapFunc func(Generic, int) Generic

func Map(elements []Generic, callback GenericMapFunc) []Generic {
	var result []Generic = make([]Generic, 0)
	for index, element := range elements {
		opOut := callback(element, index)
		result = append(result, opOut)
	}
	return result
}

type FilterFunc func(Generic, int) bool

func Filter(elements []Generic, callback FilterFunc) []Generic {
	var result []Generic = make([]Generic, 0)
	for index, element := range elements {
		opOut := callback(element, index)
		if opOut {
			result = append(result, opOut)
		}
	}
	return result
}

func Each(elements []Generic, callback FilterFunc) []Generic {
	var result []Generic = make([]Generic, 0)
	for index, element := range elements {
		opOut := callback(element, index)
		if opOut {
			result = append(result, opOut)
		}
	}
	return result
}

func Once(elements []Generic, callback FilterFunc) bool {
	result := Find(elements, callback)
	return result != nil
}

func Find(elements []Generic, callback FilterFunc) Generic {
	for index, element := range elements {
		opOut := callback(element, index)
		if opOut {
			return element
		}
	}
	return nil
}

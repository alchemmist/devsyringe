package utils

func Remove[T comparable](slice []T, value T) []T {
    result := make([]T, 0, len(slice))
    for _, v := range slice {
        if v != value {
            result = append(result, v)
        }
    }
    return result
}


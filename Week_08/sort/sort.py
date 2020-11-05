def bubble_sort(nums):
    for i in range(len(nums) - 1):
        for j in range(len(nums) - 1 - i):
            if nums[j] > nums[j + 1]:
                nums[j], nums[j + 1] = nums[j + 1], nums[j]


def selection_sort(nums):
    for i in range(len(nums) - 1):
        min_index = i
        for j in range(i, len(nums)):
            if nums[j] < nums[min_index]:
                min_index = j
        nums[i], nums[min_index] = nums[min_index], nums[i]


def insertion_sort(nums):
    for i in range(1, len(nums)):
        pre_index = i - 1
        current = nums[i]
        while pre_index >= 0 and nums[pre_index] > current:
            nums[pre_index], nums[pre_index + 1] = nums[pre_index + 1], nums[pre_index]
            pre_index -= 1
        nums[pre_index + 1] = current


def merge(left, right):
    i, j = 0, 0
    res = []

    while i < len(left) and j < len(right):
        if left[i] < right[j]:
            res.append(left[i])
            i += 1
        else:
            res.append(right[j])
            j += 1
    while i < len(left):
        res.append(left[i])
        i += 1
    while j < len(right):
        res.append(right[j])
        j += 1
    return res


def merge_sort(nums):
    if len(nums) < 2:
        return nums
    mid = len(nums) // 2
    merge(merge_sort(nums[:mid]), merge_sort(nums[mid:]))


def partition(nums):
    pivot_index = 0
    pivot = nums[pivot_index]
    j = pivot_index + 1
    for i in range(j, len(nums)):
        if nums[i] < pivot:
            nums[i], nums[j] = nums[i], nums[j]
            j += 1
    nums[j - 1], nums[pivot_index] = nums[pivot_index], nums[j - 1]
    return j - 1


def quick_sort(nums):
    if len(nums) < 2:
        return
    partition_index = partition(nums)
    quick_sort(nums[:partition_index])
    quick_sort(nums[partition_index + 1:])


def heapify(parent_index, length, nums):
    temp = nums[parent_index]
    child_index = 2 * parent_index + 1
    while child_index < length:
        if child_index + 1 < length and nums[child_index + 1] > nums[child_index]:
            child_index = child_index + 1
        if temp > nums[child_index]:
            break
        nums[parent_index] = nums[child_index]
        parent_index = child_index
        child_index = 2 * parent_index + 1
    nums[parent_index] = temp


def heap_sort(nums):
    for i in range((len(nums) - 2) // 2, -1, -1):
        heapify(i, len(nums), nums)
    for j in range(len(nums) - 1, 0, -1):
        nums[j], nums[0] = nums[0], nums[j]
        heapify(0, j, nums)


if __name__ == "__main__":
    arr = [4, 6, 5, 1, 3, 2]
    heap_sort(arr)
    print(arr)

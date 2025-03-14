func productExceptSelf(nums []int) []int {
    leftProd:=1
    rightProd:=1
    zeros := 0
    for _, num := range nums {
        if num == 0 {
            zeros++
            continue
        }
        leftProd *=num
    }
    if zeros> 1 {
        for i := 0; i < len(nums); i++ {
            nums[i] = 0
        }
        return nums
    }

    tmp := 1
    for i := len(nums)-1; i >= 0; i-- {
        if zeros > 0 {
            if nums[i] != 0 {
                nums[i] = 0
            } else {
                nums[i]=leftProd
            }

        } else {
            tmp = nums[i]
            leftProd = leftProd/tmp
            nums[i] = leftProd*rightProd
            rightProd *= tmp 
        }
    }
    return nums
}



defmodule Solution do
  @spec two_sum(nums :: [integer], target :: integer) :: [integer]
  def two_sum(nums, target) do
    Solution.find_two_sum(nums, target)
  end
  def find_two_sum([head|tail], target, seen \\ %{}, i \\ 0) do
    if Map.has_key?(seen, target-head) do
      [Map.get(seen, target-head), i]
    else
      find_two_sum(tail, target, Map.put_new(seen, head, i), i + 1)
    end
  end
end

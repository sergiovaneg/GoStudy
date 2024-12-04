const reA = r"mul\((\d{1,3}),(\d{1,3})\)"
const reB = r"(?<=do\(\)).+?(?=don't\(\))"
corruptedMemory = read("./input.txt", String)

function processMemory(mem)
  return sum(
    eachmatch(reA, mem)
  ) do match
    parse(Int, match[1]) * parse(Int, match[2])
  end
end

println(corruptedMemory |> processMemory)

function filterMemory(mem)
  return map(
    match -> match.match,
    eachmatch(reB, "do()" * mem * "don't()")
  ) |> join
end

println(corruptedMemory |> filterMemory |> processMemory)

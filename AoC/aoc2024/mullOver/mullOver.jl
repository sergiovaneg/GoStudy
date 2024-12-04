corruptedMemory = readlines("./input.txt") |> join

function processMemory(mem)
  return sum(
    match -> parse(Int, match[1]) * parse(Int, match[2]),
    eachmatch(r"mul\(([0-9]{1,3}),([0-9]{1,3})\)", mem)
  )
end

println(corruptedMemory |> processMemory)

function filterMemory(mem)
  return join(
    map(
      match -> match.match,
      eachmatch(r"(?<=do\(\)).+?(?=don't\(\))", "do()" * mem * "don't()")
    )
  )
end

println(corruptedMemory |> filterMemory |> processMemory)

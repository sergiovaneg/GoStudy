corruptedMemory = readlines("./input.txt") |> join

function processMemory(line::String)::Int
  return sum(
    match -> prod(e -> parse(Int, e), match.captures),
    eachmatch(r"mul\(([0-9]{1,3}),([0-9]{1,3})\)", line)
  )
end

println(corruptedMemory |> processMemory)

function filterMemory(line::String)::String
  return join(
    map(
      match -> match.match,
      eachmatch(r"(?<=do\(\)).+?(?=don't\(\))", "do()" * line * "don't()")
    )
  )
end

println(corruptedMemory |> filterMemory |> processMemory)

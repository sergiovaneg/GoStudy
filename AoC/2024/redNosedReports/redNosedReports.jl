levels = readlines("./input.txt")

function parseLevel(level)
  res = map(
    num -> parse(Int, num, base=10),
    split(level, " ")
  )
  return res
end

parsed_levels = map(
  parseLevel,
  levels
)

function isSafeLevel(level, tol)
  deltas = diff(level)
  idx = findfirst(
    d -> d == 0 || abs(d) > 3 || sign(d) != sign(deltas[1]),
    deltas
  )
  if isnothing(idx)
    return true
  elseif tol == 0
    return false
  else
    return isSafeLevel(
      vcat(level[1:idx-1], level[idx+1:end]),
      tol - 1
    ) || isSafeLevel(
      vcat(level[1:idx], level[idx+2:end]),
      tol - 1
    )
  end
end

function isSafeLevel(level)
  return isSafeLevel(level, 0)
end

map(
  isSafeLevel,
  parsed_levels
) |> count |> println

map(
  l -> isSafeLevel(l, 1),
  parsed_levels
) |> count |> println

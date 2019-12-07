module Accumulate

let rec notTailRecursiveAccumulate (func: 'a -> 'b) (input: 'a list): 'b list =
    match input with
    | [] -> []
    | head :: tail ->
        func head :: (notTailRecursiveAccumulate func tail)
        
let rec accumulate (func: 'a -> 'b) (input: 'a list): 'b list =
    let rec loop acc = function
        | [] -> acc
        | head :: tail -> loop (func head :: acc) tail
    loop [] input |> List.rev
   
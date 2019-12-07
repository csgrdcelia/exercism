module QueenAttack

let between (number1: int) (number2: int) (input: int) =
    input > number1 && input < number2

let create (position: int * int) =
    (fst position |> between 0 8) && (snd position |> between 0 8) 

let canAttack (queen1: int * int) (queen2: int * int) = failwith "You need to implement this function."
module Bob
open System
open System.Text.RegularExpressions

let isYelling (input: string): bool =
    Regex.Match(input, "[A-Z]").Success && not (Regex.Match(input, "[a-z]").Success)

let isFinishedBy (char: char) (input: string) =
    (input.Trim() |> Seq.last) = char

let response (input: string): string =
    match input with
    | input when input |> String.IsNullOrWhiteSpace -> "Fine. Be that way!"
    | input when input |> isYelling && input |> isFinishedBy '?' -> "Calm down, I know what I'm doing!"
    | input when input |> isFinishedBy '?' -> "Sure."
    | input when input |> isYelling -> "Whoa, chill out!"
    | _ -> "Whatever."


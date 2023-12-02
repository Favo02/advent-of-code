% Advent of code 2023 day 02.
% Part 1 is enough.

-module(day02).
-export([part1/1]).

part1(InputFile) ->
  Lines = lists:reverse(readlines(InputFile)),
  Self = self(),
  lists:map(
    fun({I, L}) -> spawn(fun() -> parse_line(Self, I, L) end) end,
    lists:enumerate(Lines)
  ),
  Received = part1_receiver(1, length(Lines), []),
  Result = sum_valid_indexes(Received),
  io:format("Part 1: ~p~n", [Result]),
  Result.

% read lines from a file "FileName"
readlines(FileName) ->
  {ok, Device} = file:open(FileName, [read]),
  try get_all_lines(Device)
    after file:close(Device)
  end.

% get all lines from device "Device"
get_all_lines(Device) ->
  case io:get_line(Device, "") of
    eof  -> [];
    Line -> [lists:sublist(Line, 9, length(Line)-9) | get_all_lines(Device)]
  end.

% parse a single line of input
parse_line(From, Index, Line) ->
  Parts = string:split(Line, ";", all),
  Results = lists:map(
    fun(P) -> reduce_booleans(parse_part(P)) end,
    Parts
  ),
  From ! {res, Index, Results},
  ok.

% parse a single part of a line
parse_part(Part) ->
  {match, Colors} = re:run(
    Part,
    "([0-9]+)\\s*(red|green|blue)",
    [global, {capture, all_but_first, list}]
  ),
  lists:map(
    fun([Count|[Col|[]]]) -> valid_checker(Col, Count) end,
    Colors
  ).

% checks if a color is below its limit
valid_checker(Col, Count) ->
  Limits = [{"red", 12}, {"green", 13}, {"blue", 14}],
  {_, Limit} = lists:keyfind(Col, 1, Limits),
  {ValCount, _} = string:to_integer(Count),
  ValCount =< Limit.

% reduce a list of booleans to a single boolean (all true to be true)
reduce_booleans(List) ->
  lists:foldr(
    fun(I, Acc) -> I and Acc end,
    true,
    List
  ).

% receive results from various actors and reorder them (index is important)
part1_receiver(Index, Tot, Acc) when Index =< Tot ->
  receive
    {res, Index, Res} -> part1_receiver(Index+1, Tot, [reduce_booleans(Res) | Acc])
  end;
part1_receiver(_, _, Acc) -> Acc.

% sum the indexes of true elements in a list
sum_valid_indexes(List) ->
  lists:sum(
    lists:map(
      fun({Index,_}) -> Index end,
      lists:filter(
        fun({_, Bool}) -> Bool == true end,
        lists:enumerate(List)
      )
    )
  ).

-module(pnparse).
-include("records.hrl").
-import(string, [tokens/2]).
-import(parse_util, [parse_qs_key/1, parse_qs_value/1]).
-export([pnparse_qs/1]).

pnparse_qs(Query) -> 
    pnparse_qs(Query, #pn{}).
pnparse_qs([], PNRecord) ->
    PNRecord;
pnparse_qs(Query, PNRecord) ->
    {Key, Rest} = parse_qs_key(Query),
    {Value, Rest1} = parse_qs_value(Rest),
    if
        Key =:= "id" ->
            [Sid, Uid, Vid, Pid, Peid] = string:tokens(Value,"."),
            PNRecord1 = PNRecord#pn{sid=Sid, uid=Uid, vid=Vid, pid=Pid, peid=Peid};
        Key =:= "stat" ->
            [Y, Ym, Vh, St] = string:tokens(Value, "."),
            PNRecord1 = PNRecord#pn{y=Y, ym=Ym, vh=Vh, st=St};
        true ->
            PNRecord1 = PNRecord
    end,
    pnparse_qs(Rest1, PNRecord1).

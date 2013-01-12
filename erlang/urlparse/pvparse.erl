-module(pvparse).
-include("records.hrl").
-import(string, [tokens/2]).
-import(parse_util, [parse_qs_key/1, parse_qs_value/1]).
-export([pvparse_qs/1]).

pvparse_qs(Query) -> 
    pvparse_qs(Query, #pv{}).
pvparse_qs([], PVRecord) ->
    PVRecord;
pvparse_qs(Query, PVRecord) ->
    {Key, Rest} = parse_qs_key(Query),
    {Value, Rest1} = parse_qs_value(Rest),
    if
        Key =:= "id" ->
            [Sid, Uid, Vid, Pid, Peid] = string:tokens(Value,"."),
            PVRecord1 = PVRecord#pv{sid=Sid, uid=Uid, vid=Vid, pid=Pid, peid=Peid};
        Key =:= "stat" ->
            [Y, Ym, Vh, St] = string:tokens(Value, "."),
            PVRecord1 = PVRecord#pv{y=Y, ym=Ym, vh=Vh, st=St};
        true ->
            PVRecord1 = PVRecord
    end,
    pvparse_qs(Rest1, PVRecord1).

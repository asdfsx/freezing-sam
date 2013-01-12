-module(osparse).
-include("records.hrl").
-import(string, [tokens/2]).
-import(parse_util, [parse_qs_key/1, parse_qs_value/1]).
-export([osparse_qs/1]).

osparse_qs(Query) -> 
    osparse_qs(Query, #os{}).
osparse_qs([], OSRecord) ->
    OSRecord;
osparse_qs(Query, OSRecord) ->
    {Key, Rest} = parse_qs_key(Query),
    {Value, Rest1} = parse_qs_value(Rest),
    if
        Key =:= "id" ->
            [Sid, Uid, Vid, Pid, Peid] = string:tokens(Value,"."),
            OSRecord1 = OSRecord#os{sid=Sid, uid=Uid, vid=Vid, pid=Pid, peid=Peid};
        Key =:= "stat" ->
            [Y, Ym, Vh, St] = string:tokens(Value, "."),
            OSRecord1 = OSRecord#os{y=Y, ym=Ym, vh=Vh, st=St};
        true ->
            OSRecord1 = OSRecord
    end,
    osparse_qs(Rest1, OSRecord1).

-module(ocparse).
-include("records.hrl").
-import(string, [tokens/2]).
-import(parse_util, [parse_qs_key/1, parse_qs_value/1]).
-export([ocparse_qs/1]).

ocparse_qs(Query) -> 
    ocparse_qs(Query, #oc{}).
ocparse_qs([], OCRecord) ->
    OCRecord;
ocparse_qs(Query, OCRecord) ->
    {Key, Rest} = parse_qs_key(Query),
    {Value, Rest1} = parse_qs_value(Rest),
    if
        Key =:= "id" ->
            [Sid, Uid, Vid, Pid, Peid] = string:tokens(Value,"."),
            OCRecord1 = OCRecord#oc{sid=Sid, uid=Uid, vid=Vid, pid=Pid, peid=Peid};
        Key =:= "stat" ->
            [Y, Ym, Vh, St] = string:tokens(Value, "."),
            OCRecord1 = OCRecord#oc{y=Y, ym=Ym, vh=Vh, st=St};
        true ->
            OCRecord1 = OCRecord
    end,
    ocparse_qs(Rest1, OCRecord1).

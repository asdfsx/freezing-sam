-module(hbparse).
-include("records.hrl").
-import(string, [tokens/2]).
-import(parse_util, [parse_qs_key/1, parse_qs_value/1]).
-export([hbparse_qs/1]).

hbparse_qs(Query) -> 
    hbparse_qs(Query, #hb{}).
hbparse_qs([], HBRecord) ->
    HBRecord;
hbparse_qs(Query, HBRecord) ->
    {Key, Rest} = parse_qs_key(Query),
    {Value, Rest1} = parse_qs_value(Rest),
    if
        Key =:= "id" ->
            [Sid, Uid, Vid, Pid, Peid] = string:tokens(Value,"."),
            HBRecord1 = HBRecord#hb{sid=Sid, uid=Uid, vid=Vid, pid=Pid, peid=Peid};
        Key =:= "stat" ->
            [Y, Ym, Vh, St] = string:tokens(Value, "."),
            HBRecord1 = HBRecord#hb{y=Y, ym=Ym, vh=Vh, st=St};
        true ->
            HBRecord1 = HBRecord
    % 	HBRecord1 = HBRecord#hb{tt=Value}
    end,
    hbparse_qs(Rest1, HBRecord1).

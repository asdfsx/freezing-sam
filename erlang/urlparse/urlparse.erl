-module(urlparse).
-include("records.hrl").
-export([urlparse/1, testhb/0]).
-import(hbparse, [hbparse_qs/1]).
-import(osparse, [osparse_qs/1]).
-import(ocparse, [ocparse_qs/1]).
-import(pvparse, [pvparse_qs/1]).
-import(pnparse, [pnparse_qs/1]).

urlparse("/pn?" ++ Query) ->
    pnparse_qs(Query);
urlparse("/pv?" ++ Query) ->
    pvparse_qs(Query);
urlparse("/oc?" ++ Query) ->
    ocparse_qs(Query);
urlparse("/os?" ++ Query) ->
    osparse_qs(Query);
urlparse("/hb?" ++ Query) ->
    hbparse:hbparse_qs(Query);
urlparse(_) ->
    "".

testhb() ->
    urlparse("/os?id=56fbce4e.tksog8re3ZqKPIv1clNJtA.-dTAqKs-DAx/g46GRVaIhA.X8CqYXTI5zc6pT8GxUuGDg.UrjZ3y7oW8M/G/36u8kSBQ&stat=-1.-899.899.10&ptif=1&v=1.1&ts=1357401854569").


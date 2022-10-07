module Stargazing exposing (main)

import Html exposing (..)
import Html.Attributes exposing (class, src)
import Html.Events exposing (onClick)
import Browser
type Msg =
    Like | Unlike
initModel : { liked : Bool }
initModel = 
    {
        liked = True
    }
starmodel : { liked : Bool } -> Html Msg
starmodel model =
    let
        buttonType = 
            if model.liked then "star" else "star_outline"
        msg =
            if model.liked then Unlike else Like
    in
    div [][
            span [ class "material-icons md-100", onClick msg ] [ text buttonType ] ]
webview : { liked : Bool } -> Html Msg
webview model =
    starmodel model
update : Msg -> { liked : Bool } -> { liked : Bool }
update msg model =
    case msg of
        Like ->
            { model | liked = True }
        Unlike ->
            { model | liked = False }

main : Program () { liked : Bool } Msg
main =
    Browser.sandbox
    {
        init = initModel
        ,view = webview
        ,update = update
    }
 
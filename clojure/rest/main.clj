(ns rest
  (:require [compojure.core :refer :all]
            [compojure.route :as route]))


(defroutes app
  (api
   (GET "/" []
     (ok {:message (str "hello world")}))
   (GET "/hello" []
     :query-params [name :- String]
     (ok {:message (str "Hello, " name)}))))
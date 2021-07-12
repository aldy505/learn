(defproject rest "0.0.0"
  :description "Trial rest api"
  :dependencies [[compojure "1.6.2"]]
  :plugins [[lein-ring "0.12.5"]]
  :ring {:handler rest.app}
  )
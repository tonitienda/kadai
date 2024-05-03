(ns kadai.handler
  (:require [compojure.core :refer :all]
            [compojure.route :as route]
            [clojure.data.json :as json]
            [ring.middleware.defaults :refer [wrap-defaults site-defaults]]))

(defroutes app-routes
  (GET "/" [] "Hello World")
  (GET "/healthz" [] (json/write-str {:message "ok"}))
  (route/not-found "Not Found"))

(def app
  (wrap-defaults app-routes site-defaults))

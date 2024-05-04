(ns kadai.core
  (:require [org.httpkit.server :as hk-server]))

(defn app [_req]
  {:status  200
  :headers  {"Content-Type" "text/html"}
  :body "hello HTTP!"})

(defn start-server [_argc]
(hk-server/run-server app {:port 8080}))
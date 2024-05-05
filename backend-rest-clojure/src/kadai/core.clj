(ns kadai.core
  (:require [compojure.core :refer [defroutes GET POST DELETE]]
            [compojure.route :as route]
            [org.httpkit.server :as hk-server]
            [clojure.data.json :as json]
            [kadai.tasks :as tasks]
            [kadai.in-memory :as in-memory]
            ))

(defn handle-health [req]
    {:status  200
    :headers {"Content-Type" "application/json"}
    :body   (json/write-str {:message (str "ok")})}
  )

(defn make-handler [db f]
  (fn [req]
    (f db req)))


(defroutes all-routes
  (let [db (in-memory/initialize-db)]
    (GET "/healthz" [] handle-health)
    (GET "/v0/tasks" [] (make-handler db tasks/handle-get-tasks))
    (POST "/v0/tasks" [] (make-handler db tasks/handle-post-tasks))
    (DELETE "/v0/tasks/:id" [id] (make-handler db tasks/handle-delete-tasks))
    (POST "/v0/actions/undo-delete" [] (make-handler db tasks/handle-undo-delete-task))))


(defn start-server [_argc]
  (hk-server/run-server all-routes {:port 8080}))
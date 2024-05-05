(ns kadai.tasks
  (:require [clojure.data.json :as json]))


(defprotocol TasksDataSource
  (get-tasks [this owner-id])
  (get-task [this id])
  (add-task [this task])
  (delete-task [this id])
  (update-task [this task]))


(defn handle-get-tasks [db req]
    {:status  200
    :headers {"Content-Type" "application/json"}
    :body   (json/write-str (get-tasks db "03c6eb30-e65f-49fd-92c2-a207f03bbf51"))}
)


(defn handle-post-tasks [db req]
    {:status  202
    :headers {"Content-Type" "application/json"}
    :body   nil}
)

(defn handle-undo-delete-task [db req]
    {:status  202
    :headers {"Content-Type" "application/json"}
    :body   nil}
)


(defn handle-delete-tasks [db req]
    {:status  200
    :headers {"Content-Type" "application/json"}
    :body   (json/write-str {
      :url "/v0/tasks/12345/undo-delete"
      :method "POST"
    })}
)

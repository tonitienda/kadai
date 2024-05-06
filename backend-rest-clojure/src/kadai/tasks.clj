(ns kadai.tasks
  (:require 
    [clojure.data.json :as json]))

(require '[malli.core :as m])

(def TaskRequest 
    [:map 
        [:title :string ]
        [:description :string ]
    ])

(defprotocol TasksDataSource
  (get-tasks [this owner-id])
  (get-task [this id])
  (add-task [this task])
  (delete-task [this id])
  (update-task [this task]))



(defn handle-get-tasks [db req]
    (let [
        owner-id  "03c6eb30-e65f-49fd-92c2-a207f03bbf51"
        body (json/write-str (get-tasks db owner-id))]
    {:status  200
    :headers {"Content-Type" "application/json"}
    :body   body})
)

(defn- new-task [title description owner-id]
    {
        :id "8ea6c732-4d15-44a6-8864-0b5e0a7042fd"
        :title title
        :description description
        :status "pending"
        :owner-id owner-id
    })

(defn handle-post-tasks [db req]
  (let [body (:body req)]
    (try (m/validate TaskRequest body)
         (let [
            title (:title body)
            description (:description body)
            user-id "66f2ce13-d1e7-4b06-ab5a-c5e7c6a55fcf" ;; get user id from header
            task (new-task title description user-id)] 
         (add-task db task)
         {:status 200
          :headers {"Content-Type" "application/json"}
          :body task})
         (catch Exception e
           {:status 400
            :body {:message (.getMessage e)}}))))

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

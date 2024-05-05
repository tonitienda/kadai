(ns kadai.in-memory
  (:require [datomic.api :as d]
            [kadai.tasks :as tasks]))


(defrecord DatomicTasksDataSource [db]
  tasks/TasksDataSource
  (get-tasks [this owner-id]
  (let [query '[:find ?title ?description ?status ?deleted-at ?deleted-by
                :where
                [?task :task/owner ?owner]
                [?task :task/title ?title]
                [?task :task/description ?description]
                [?task :task/status ?status]
                [?task :task/deleted-at ?deleted-at]
                [?task :task/deleted-by ?deleted-by]
                [(= ?owner owner-id)]]
        results (d/q query db)]
    (map (fn [[title description status deleted-at deleted-by]]
           {:owner owner-id
            :title title
            :description description
            :status status
            :deleted-at deleted-at
            :deleted-by deleted-by})
         results)))

  (get-task [this id]
    (let [query '[:find ?owner ?title ?description ?status ?deleted-at ?deleted-by
                  :where
                  [?task :task/id ?id]
                  [?task :task/owner ?owner]
                  [?task :task/title ?title]
                  [?task :task/description ?description]
                  [?task :task/status ?status]
                  [?task :task/deleted-at ?deleted-at]
                  [?task :task/deleted-by ?deleted-by]]
          results (d/q query db)]
      (if (seq results)
        (let [[owner title description status deleted-at deleted-by] (first results)]
          {:id id
           :owner owner
           :title title
           :description description
           :status status
           :deleted-at deleted-at
           :deleted-by deleted-by})
        nil)))

  (add-task [this task]
    (d/transact db `[(-> :task
                          (d/tempid :db.part/user)
                          :task/id ~(-> task :id)
                          :task/owner ~(-> task :owner-id)
                          :task/title ~(-> task :title)
                          :task/description ~(-> task :description)
                          :task/status ~(-> task :status)
                          :task/deleted-at ~(-> task :deleted-at)
                          :task/deleted-by ~(-> task :deleted-by))]))

  (delete-task [this id]
    (d/transact db `[(d/retract db [:task/id ~id])]))

  (update-task [this task]
    (d/transact db `[(d/retract db [:task/id ~(-> task :id)])
                     (-> task
                         (update-in [:deleted-at] #(.toMillis %))
                         (d/tempid :db.part/user)
                         :task/id ~(-> task :id)
                         :task/owner ~(-> task :owner-id)
                         :task/title ~(-> task :title)
                         :task/description ~(-> task :description)
                         :task/status ~(-> task :status)
                         :task/deleted-at ~(-> task :deleted-at)
                         :task/deleted-by ~(-> task :deleted-by))])))



(defn initialize-db []
  (let [uri "datomic:mem://kadai"
        db (d/create-database uri)
        conn (d/connect uri)]
    (->DatomicTasksDataSource conn)))
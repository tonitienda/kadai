(ns kadai.handler
  (:require [compojure.core :refer :all]
            [compojure.route :as route]
            [clojure.data.json :as json]
            [ring.middleware.json :as json-middleware]
            [ring.util.request :as request]
            [ring.util.response :refer [response]]
            [ring.middleware.defaults :refer [wrap-defaults site-defaults]]))

(defn empty-task-list []
  [])

(defn- empty-response []
  (response {:status 200 :headers {"Content-Type" "application/json"}}))

(defn log-request-before [handler]
  (fn [request]
    (let [response (handler request)]
      (println (str "Request: " (:request-method request) " " (:uri request)))
      response)))

(defn log-request [handler]
  (fn [request]
    (let [response (handler request)]
      (println (str "Request: " (:request-method request) " " (:uri request) " -> " (:status response)))
      response)))

(defroutes app-routes
  (GET "/" [] "Hello World")
  (GET "/healthz" [] (json/write-str {:message "ok"}))

  (GET "/v0/tasks" [] (json-middleware/wrap-json-response (response (empty-task-list))))
  (POST "/v0/tasks" [title description]
    (let [task {:title title :description description}]
      (empty-response)))
  (DELETE "/v0/tasks/:id" [] (json-middleware/wrap-json-response(json/write-str {:something "something"})))

  (route/not-found "Not Found"))

;; TODO - Temporarily disable the antiforgery check
;; I will enable it and add it to the rest of backends later on
(def app  (-> app-routes
  (log-request-before)
  (wrap-defaults (assoc-in site-defaults [:security :anti-forgery] false))
  (log-request)))

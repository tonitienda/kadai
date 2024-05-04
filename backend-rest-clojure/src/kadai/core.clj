(ns kadai.core
  (:require [compojure.core :refer [defroutes GET POST]]
            [compojure.route :as route]
            [org.httpkit.server :as hk-server]
            [clojure.data.json :as json]))

(defn handle-health [req]
    {:status  200
    :headers {"Content-Type" "application/json"}
    :body   (json/write-str {:message "ok"})}
  )

(defroutes all-routes
  (GET "/healthz" [] handle-health)
  ;;(GET "/ws" [] chat-handler)     ;; websocket
  ;;(GET "/async" [] async-handler) ;; asynchronous(long polling)
  ;; (context "/user/:id" []
  ;;          (GET / [] get-user-by-id)
  ;;          (POST / [] update-userinfo))
  ;; (files "/static/") ;; static file url prefix /static, in `public` folder
  ;;(not-found "<p>Page not found.</p>")) ;; all other, return 404
)

(defn start-server [_argc]
(hk-server/run-server all-routes {:port 8080}))
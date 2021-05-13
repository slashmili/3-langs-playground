(ns github-starts.core
  (:require
   [clj-http.client :as client]
   [clojure.data.json :as json]
   [clojure.string :as str]
   [slingshot.slingshot :as st]))

(def request-uri (format "https://api.github.com/repos/%s/%s/stargazers" "rust-lang-nursery" "rust-cookbook"))

(defn fetch [url]
  (st/try+
   {:ok (client/get url)}
   (catch [:status 404] {:keys [request-time headers body]}
     {:error, :not_found})
   (catch Object _
     (st/throw+))))
(defn fetch-users [] (fetch request-uri))

(defn starts-with-s? [string] str/starts-with? string "s")

(defn parse-and-filter [response]
  (->> response
       (:body)
       (json/read-str)
       (filter starts-with-s?)))

(defn run-sync []
  (let [response (:ok (fetch-users))]
    (if response
      (println (parse-and-filter response))
      "Explosion Detected in HTTP Realm!")))

(defn run-async []
  (client/get request-uri
            {:async? true}
            ;; respond callback
            (fn [response] (println (parse-and-filter response)))
            ;; raise callback
            (fn [exception] (println "Explosion Detected in HTTP Realm!"))))

(defn -main
  [& args]
  (run-async)
)

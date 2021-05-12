(ns github-starts.core
  (:gen-class)
  (:require [clj-http.client :as client] [clojure.data.json :as json] [clojure.string :as str])
  )

(def request-uri (format "https://api.github.com/repos/%s/%s/stargazers" "rust-lang-nursery" "rust-cookbook"))

(defn fetch-users
  []
  (client/get request-uri))
(defn starts-with-s?
  [string]
  str/starts-with? string "s"
  )
(defn -main
  [& args]
  (->>
  (fetch-users)
  (:body)
  (json/read-str)
  (filter starts-with-s?)
  (println)
  )
)

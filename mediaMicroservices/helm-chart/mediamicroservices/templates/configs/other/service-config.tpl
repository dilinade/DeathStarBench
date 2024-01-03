{{- define "mediamicroservices.templates.other.service-config.json"  }}
{
  "secret": "secret",
  "unique-id-service": {
    "addr": "unique-id-service-knative.default.svc.cluster.local",
    "port": 80
  },
  "movie-id-service": {
    "addr": "movie-id-service-knative.default.svc.cluster.local",
    "port": 80
  },
  "movie-id-mongodb": {
    "addr": "movie-id-mongodb",
    "port": 27017
  },
  "movie-id-memcached": {
    "addr": "movie-id-memcached",
    "port": 11211
  },
  "user-mongodb": {
    "addr": "user-mongodb",
    "port": 27017
  },
  "user-memcached": {
    "addr": "user-memcached",
    "port": 11211
  },
  "text-service": {
    "addr": "text-service-knative.default.svc.cluster.local",
    "port": 80
  },
  "rating-service": {
    "addr": "rating-service-knative.default.svc.cluster.local",
    "port": 80
  },
  "rating-redis": {
    "addr": "rating-redis",
    "port": 6379
  },
  "user-service": {
    "addr": "user-service-knative.default.svc.cluster.local",
    "port": 80
  },
  "compose-review-service": {
    "addr": "compose-review-service-knative.default.svc.cluster.local",
    "port": 80
  },
  "compose-review-memcached": {
    "addr": "compose-review-memcached",
    "port": 11211
  },
  "review-storage-service": {
    "addr": "review-storage-service-knative.default.svc.cluster.local",
    "port": 80
  },
  "review-storage-mongodb": {
    "addr": "review-storage-mongodb",
    "port": 27017
  },
  "review-storage-memcached": {
    "addr": "review-storage-memcached",
    "port": 11211
  },
  "user-review-service": {
    "addr": "user-review-service-knative.default.svc.cluster.local",
    "port": 80
  },
  "user-review-mongodb": {
    "addr": "user-review-mongodb",
    "port": 27017
  },
  "user-review-redis": {
    "addr": "user-review-redis",
    "port": 6379
  },
  "movie-review-service": {
    "addr": "movie-review-service-knative.default.svc.cluster.local",
    "port": 80
  },
  "movie-review-mongodb": {
    "addr": "movie-review-mongodb",
    "port": 27017
  },
  "movie-review-redis": {
    "addr": "movie-review-redis",
    "port": 6379
  },
  "cast-info-service": {
    "addr": "cast-info-service-knative.default.svc.cluster.local",
    "port": 80
  },
  "cast-info-mongodb": {
    "addr": "cast-info-mongodb",
    "port": 27017
  },
  "cast-info-memcached": {
    "addr": "cast-info-memcached",
    "port": 11211
  },
  "plot-service": {
    "addr": "plot-service-knative.default.svc.cluster.local",
    "port": 80
  },
  "plot-mongodb": {
    "addr": "plot-mongodb",
    "port": 27017
  },
  "plot-memcached": {
    "addr": "plot-memcached",
    "port": 11211
  },
  "movie-info-service": {
    "addr": "movie-info-service-knative.default.svc.cluster.local",
    "port": 80
  },
  "movie-info-mongodb": {
    "addr": "movie-info-mongodb",
    "port": 27017
  },
  "movie-info-memcached": {
    "addr": "movie-info-memcached",
    "port": 11211
  },
  "page-service": {
    "addr": "page-service-knative.default.svc.cluster.local",
    "port": 80
  }
}
{{- end }}
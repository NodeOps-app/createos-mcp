using Workerd = import "/workerd/workerd.capnp";

const config :Workerd.Config = (
  services = [
    (name = "dispatcher", worker = .dispatcher),
    (name = "api-proxy",  worker = .apiProxy),
    (name = "internet", network = (allow = ["public", "private"])),
  ],
  sockets = [
    (name = "http", address = "127.0.0.1:8787", http = (), service = "dispatcher"),
  ],
);

const dispatcher :Workerd.Worker = (
  modules = [
    (name = "dispatcher.js",  esModule = embed "dispatcher.js"),
    (name = "spec-loader.js", esModule = embed "dist/spec-loader.bundled.js"),
    (name = "host.js",        esModule = embed "host.js"),
  ],
  bindings = [
    (name = "API",    service = "api-proxy"),
    (name = "LOADER", workerLoader = (id = "main")),
  ],
  compatibilityDate = "2024-09-09",
  compatibilityFlags = ["nodejs_compat"],
);

const apiProxy :Workerd.Worker = (
  modules = [
    (name = "api-proxy.js", esModule = embed "api-proxy.js"),
  ],
  bindings = [
    (name = "BACKEND_URL", fromEnvironment = "BACKEND_URL"),
  ],
  globalOutbound = "internet",
  compatibilityDate = "2024-09-09",
);

HOST = 'http://localhost:3000'

angular.module('app', [
  'ngResource',
  'ngRoute',
  'app.controllers',
  'app.models'
])

.value('resourceUrlPrefix', HOST)

.config(['$routeProvider',
  function($routeProvider) {
    $routeProvider
      .when('/apps', {
        templateUrl: 'partials/app-list.html',
        controller: 'AppListCtrl'
      })
      .when('/apps/create', {
        templateUrl: 'partials/create-app.html',
        controller: 'CreateAppCtrl'
      })
      .when('/apps/:app_id', {
        templateUrl: 'partials/app-details.html',
        controller: 'AppDetailsCtrl'
      })
      .otherwise({
        rediretTo: '/apps'
      })
  }
])

.run(['$http',
  function($http) {
    $http.defaults.useXDomain = true
    delete $http.defaults.headers.common['X-Requested-With']
    $http.defaults.headers.common['Access-Control-Allow-Origin'] = '*'
    $http.defaults.headers.common['Access-Control-Allow-Methods'] = 'POST, GET, OPTIONS, PUT'

    $http.defaults.headers.common['Accept'] = 'application/json'
    $http.defaults.headers.common['Content-Type'] = 'application/json'

    $http.defaults.headers.common.Authorization = 'Basic YWRtaW46Z3Vlc3NtZQ=='
  }
])

// 自动给请求的url加前缀（$resource）
// .config(['$provide',
//   function($provide) {
//     return $provide.decorator('$resource', ['$delegate', 'resourceUrlPrefix',
//       function($resource, urlPrefix) {
//         debugger
//         var addPrefix = function(url) {
//           if (!urlPrefix) {
//             return url
//           }
//           if (/^http:\/\//.test(url)) {
//             return url
//           }
//           return url.replace(/^\//, urlPrefix + '/')
//         }
//         return function(url, paramDefaults, actions) {
//           url = addPrefix(url)
//           if (!actions) {
//             return $resource(url, paramDefaults, actions)
//           }
//           actions = angular.copy(actions)
//           _(actions).forEach(function(opt, action) {
//             if (opt.url != null) {
//               return opt.url = addPrefix(opt.url)
//             }
//           })
//           return $resource(url, paramDefaults, actions)
//         }
//       }
//     ])
//   }
// ])

.config([
  '$provide',
  function($provide) {
    return $provide.decorator('$http', [
      '$delegate', '$q', 'resourceUrlPrefix',
      function($http, $q, urlPrefix) {
        var createShortMethods, createShortMethodsWithData, makePromiseLike$resource, replacement;
        makePromiseLike$resource = function(promise, config) {
          return promise.then(function(resp) {
            var _ref;
            if ((_ref = config.callbacks) != null) {
              if (typeof _ref.success === "function") {
                _ref.success(resp.data, resp.headers);
              }
            }
            return resp.data;
          }, function(resp) {
            var _ref;
            if ((_ref = config.callbacks) != null) {
              if (typeof _ref.error === "function") {
                _ref.error(resp);
              }
            }
            return $q.reject(resp);
          });
        };
        createShortMethods = function() {
          return angular.forEach(arguments, function(method) {
            return replacement[method] = function(url, config) {
              return replacement(angular.extend(config || {}, {
                method: method,
                url: url
              }));
            };
          });
        };
        createShortMethodsWithData = function() {
          return angular.forEach(arguments, function(method) {
            return replacement[method] = function(url, data, config) {
              return replacement(angular.extend(config || {}, {
                method: method,
                url: url,
                data: data
              }));
            };
          });
        };
        replacement = function(requestConfig) {
          var config, startWithSlashRE;
          if (!requestConfig.requestLike$resource) {
            return $http(requestConfig);
          }
          startWithSlashRE = /^\//;
          config = _.clone(requestConfig);
          if (startWithSlashRE.test(config.url)) {
            config.url = config.url.replace(startWithSlashRE, urlPrefix + '/');
          }
          return makePromiseLike$resource($http(config), config);
        };
        createShortMethods('get', 'delete', 'head', 'jsonp');
        createShortMethodsWithData('post', 'put');
        replacement.defaults = $http.defaults;
        return replacement;
      }
    ]);
  }
])
angular.module('app.models', [])

.factory('App', ['$resource',
  function($resource) {
    var App = $resource(
      '/api/app/:app_id', {
        app_id: '@app_id'
      }, {
        'update': {
          method: 'PUT'
        }
      }
    )
    return App;
  }
])

.factory('Version', ['$resource',
  function($resource) {
    var Version = $resource('/api/app/:app_id/version/:version_id', {
      app_id: '@app_id',
      version_id: '@version_id'
    }, {
      'update': {
        method: 'PUT'
      }
    })
    return Version
  }
])

.factory('Channel', ['$resource',
  function($resource) {
    var Channel = $resource('/api/app/:app_id/channel/:channel_id', {
      app_id: '@app_id',
      channel_id: '@channel_id'
    }, {
      'update': {
        method: 'PUT'
      }
    })
    return Channel
  }
])
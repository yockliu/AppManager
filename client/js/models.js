angular.module('app.models', [])

.factory('App', ['$resource',
  function($resource) {
    var App = $resource(
      'http://localhost:3000/api/app/:app_id', {
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
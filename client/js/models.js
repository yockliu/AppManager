angular.module('app.models', [])

.factory('App', ['$resource',
  function($resource) {
    var App = $resource(
      'http://localhost:3000/api/app/:app_id', 
      {app_id: '@app_id'},
      {
        'update': {
          method: 'PUT'
        }
      }
    )
    return App;
  }
])
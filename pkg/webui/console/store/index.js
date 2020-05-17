// Copyright © 2019 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import { mapValues } from 'lodash'
import * as Sentry from '@sentry/browser'
import { createStore, applyMiddleware, compose } from 'redux'
import { createLogicMiddleware } from 'redux-logic'
import { routerMiddleware } from 'connected-react-router'
import createSentryMiddleware from 'redux-sentry-middleware'

import dev from '@ttn-lw/lib/dev'
import env from '@ttn-lw/lib/env'

import createRootReducer from './reducers'
import requestPromiseMiddleware from './middleware/request-promise-middleware'
import logics from './middleware/logics'

if (env.sentryDsn)
  Sentry.init({
    dsn: env.sentryDsn,
    release: process.env.VERSION,
    normalizeDepth: 10,
  })

const composeEnhancers = (dev && window.__REDUX_DEVTOOLS_EXTENSION_COMPOSE__) || compose

export default function(history) {
  const middleware = applyMiddleware(
    createSentryMiddleware(Sentry, {
      actionTransformer: action => {
        switch (action.type) {
          case 'GET_API_KEY_LIST_SUCCESS':
          case 'GET_API_KEY_SUCCESS':
            return {
              ...action,
              payload: {
                ...action.payload,
                entities: undefined,
              },
            }
          default:
            return action
        }
      },
      stateTransformer: state => {
        return {
          ...state,
          apiKeys: undefined,
          collaborators: {
            entities: {
              ...state.collaborators.entities,
              admin: {
                ...state.collaborators.entities.admin,
                ids: undefined,
              },
            },
          },
          devices: {
            ...state.devices,
            entities: mapValues(state.devices.entities, value => {
              if (Boolean(value.ids)) {
                return { ...value, ids: undefined }
              }
            }),
          },
          gateways: {
            ...state.gateways,
            entities: mapValues(state.gateways.entities, value => {
              if (Boolean(value.ids)) {
                return { ...value, ids: undefined }
              }
            }),
          },
          pagination: undefined,
          user: {
            ...state.user,
            user: {
              ...state.user.user,
              ids: undefined,
            },
          },
          users: undefined,
        }
      },
    }),
    requestPromiseMiddleware,
    routerMiddleware(history),
    createLogicMiddleware(logics),
  )

  const store = createStore(createRootReducer(history), composeEnhancers(middleware))
  if (dev && module.hot) {
    module.hot.accept('./reducers', () => {
      store.replaceReducer(createRootReducer(history))
    })
  }

  return store
}

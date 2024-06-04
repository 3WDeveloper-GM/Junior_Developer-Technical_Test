/* prettier-ignore-start */

/* eslint-disable */

// @ts-nocheck

// noinspection JSUnusedGlobalSymbols

// This file is auto-generated by TanStack Router

// Import Routes

import { Route as rootRoute } from './routes/__root'
import { Route as LoginImport } from './routes/login'
import { Route as AboutImport } from './routes/about'
import { Route as BillsImport } from './routes/_bills'
import { Route as IndexImport } from './routes/index'
import { Route as BillsBillsUpdateImport } from './routes/_bills/bills-update'
import { Route as BillsBillsRemoveImport } from './routes/_bills/bills-remove'
import { Route as BillsBillsReadImport } from './routes/_bills/bills-read'
import { Route as BillsBillsFetchImport } from './routes/_bills/bills-fetch'
import { Route as BillsBillsCreateImport } from './routes/_bills/bills-create'

// Create/Update Routes

const LoginRoute = LoginImport.update({
  path: '/login',
  getParentRoute: () => rootRoute,
} as any)

const AboutRoute = AboutImport.update({
  path: '/about',
  getParentRoute: () => rootRoute,
} as any)

const BillsRoute = BillsImport.update({
  id: '/_bills',
  getParentRoute: () => rootRoute,
} as any)

const IndexRoute = IndexImport.update({
  path: '/',
  getParentRoute: () => rootRoute,
} as any)

const BillsBillsUpdateRoute = BillsBillsUpdateImport.update({
  path: '/bills-update',
  getParentRoute: () => BillsRoute,
} as any)

const BillsBillsRemoveRoute = BillsBillsRemoveImport.update({
  path: '/bills-remove',
  getParentRoute: () => BillsRoute,
} as any)

const BillsBillsReadRoute = BillsBillsReadImport.update({
  path: '/bills-read',
  getParentRoute: () => BillsRoute,
} as any)

const BillsBillsFetchRoute = BillsBillsFetchImport.update({
  path: '/bills-fetch',
  getParentRoute: () => BillsRoute,
} as any)

const BillsBillsCreateRoute = BillsBillsCreateImport.update({
  path: '/bills-create',
  getParentRoute: () => BillsRoute,
} as any)

// Populate the FileRoutesByPath interface

declare module '@tanstack/react-router' {
  interface FileRoutesByPath {
    '/': {
      id: '/'
      path: '/'
      fullPath: '/'
      preLoaderRoute: typeof IndexImport
      parentRoute: typeof rootRoute
    }
    '/_bills': {
      id: '/_bills'
      path: ''
      fullPath: ''
      preLoaderRoute: typeof BillsImport
      parentRoute: typeof rootRoute
    }
    '/about': {
      id: '/about'
      path: '/about'
      fullPath: '/about'
      preLoaderRoute: typeof AboutImport
      parentRoute: typeof rootRoute
    }
    '/login': {
      id: '/login'
      path: '/login'
      fullPath: '/login'
      preLoaderRoute: typeof LoginImport
      parentRoute: typeof rootRoute
    }
    '/_bills/bills-create': {
      id: '/_bills/bills-create'
      path: '/bills-create'
      fullPath: '/bills-create'
      preLoaderRoute: typeof BillsBillsCreateImport
      parentRoute: typeof BillsImport
    }
    '/_bills/bills-fetch': {
      id: '/_bills/bills-fetch'
      path: '/bills-fetch'
      fullPath: '/bills-fetch'
      preLoaderRoute: typeof BillsBillsFetchImport
      parentRoute: typeof BillsImport
    }
    '/_bills/bills-read': {
      id: '/_bills/bills-read'
      path: '/bills-read'
      fullPath: '/bills-read'
      preLoaderRoute: typeof BillsBillsReadImport
      parentRoute: typeof BillsImport
    }
    '/_bills/bills-remove': {
      id: '/_bills/bills-remove'
      path: '/bills-remove'
      fullPath: '/bills-remove'
      preLoaderRoute: typeof BillsBillsRemoveImport
      parentRoute: typeof BillsImport
    }
    '/_bills/bills-update': {
      id: '/_bills/bills-update'
      path: '/bills-update'
      fullPath: '/bills-update'
      preLoaderRoute: typeof BillsBillsUpdateImport
      parentRoute: typeof BillsImport
    }
  }
}

// Create and export the route tree

export const routeTree = rootRoute.addChildren({
  IndexRoute,
  BillsRoute: BillsRoute.addChildren({
    BillsBillsCreateRoute,
    BillsBillsFetchRoute,
    BillsBillsReadRoute,
    BillsBillsRemoveRoute,
    BillsBillsUpdateRoute,
  }),
  AboutRoute,
  LoginRoute,
})

/* prettier-ignore-end */

/* ROUTE_MANIFEST_START
{
  "routes": {
    "__root__": {
      "filePath": "__root.tsx",
      "children": [
        "/",
        "/_bills",
        "/about",
        "/login"
      ]
    },
    "/": {
      "filePath": "index.tsx"
    },
    "/_bills": {
      "filePath": "_bills.tsx",
      "children": [
        "/_bills/bills-create",
        "/_bills/bills-fetch",
        "/_bills/bills-read",
        "/_bills/bills-remove",
        "/_bills/bills-update"
      ]
    },
    "/about": {
      "filePath": "about.tsx"
    },
    "/login": {
      "filePath": "login.tsx"
    },
    "/_bills/bills-create": {
      "filePath": "_bills/bills-create.tsx",
      "parent": "/_bills"
    },
    "/_bills/bills-fetch": {
      "filePath": "_bills/bills-fetch.tsx",
      "parent": "/_bills"
    },
    "/_bills/bills-read": {
      "filePath": "_bills/bills-read.tsx",
      "parent": "/_bills"
    },
    "/_bills/bills-remove": {
      "filePath": "_bills/bills-remove.tsx",
      "parent": "/_bills"
    },
    "/_bills/bills-update": {
      "filePath": "_bills/bills-update.tsx",
      "parent": "/_bills"
    }
  }
}
ROUTE_MANIFEST_END */

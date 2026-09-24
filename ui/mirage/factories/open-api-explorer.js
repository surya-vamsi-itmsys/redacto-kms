/**
 * Copyright (c) HashiCorp, Inc.
 * SPDX-License-Identifier: MPL-2.0
 */

import { Factory } from 'miragejs';
/* eslint-disable ember/avoid-leaking-state-in-ember-objects */
export default Factory.extend({
  openapi: '3.0.2',
  info: {
    title: 'Redacto KMS API',
    description:
      'HTTP API that gives you full access to Redacto KMS. All API routes are prefixed with `/v1/`.',
    version: '1.0.0',
    license: {
      name: '*',
    },
  },
  paths: {
    '/auth/token/create': {
      description: 'The token create path is used to create new tokens.',
      post: {
        summary: 'The token create path is used to create new tokens.',
        tags: ['auth'],
        responses: {
          200: {
            description: 'OK',
          },
        },
      },
    },
    '/secret/data/{path}': {
      description: 'Location of a secret.',
      post: {
        summary: 'Location of a secret.',
        tags: ['secret'],
        responses: {
          200: {
            description: 'OK',
          },
        },
      },
    },
  },
});

// Copyright 2022 The Kubermatic Kubernetes Platform contributors.
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

import {Endpoint} from '../../../../utils/endpoint';
import {SSHKeyStrategy} from './types';

export class MockedSSHKeyStrategy implements SSHKeyStrategy {
  private static readonly _fixturePath = 'ssh-keys.json';
  private static readonly _fixtureEmptyArrayPath = 'empty-array.json';
  private static _activeFixture = MockedSSHKeyStrategy._fixtureEmptyArrayPath;

  constructor() {
    this._init();
  }

  onCreate(): void {
    MockedSSHKeyStrategy._activeFixture = MockedSSHKeyStrategy._fixturePath;
  }

  onDelete(): void {
    MockedSSHKeyStrategy._activeFixture = MockedSSHKeyStrategy._fixtureEmptyArrayPath;
  }

  private _init(): void {
    cy.intercept(Endpoint.SSHKeys, req => {
      req.reply({fixture: MockedSSHKeyStrategy._activeFixture});
    });
  }
}

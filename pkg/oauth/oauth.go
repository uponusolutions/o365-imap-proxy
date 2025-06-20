// Copyright 2022 Linka Cloud  All rights reserved.
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

package oauth

import (
	"context"

	"github.com/AzureAD/microsoft-authentication-library-for-go/apps/confidential"
)

type Provider struct {
	client confidential.Client
}

func (p *Provider) GetAccessToken() (string, error) {
	scopes := []string{"https://outlook.office365.com/.default"}
	result, err := p.client.AcquireTokenSilent(context.TODO(), scopes)
	if err != nil {
		// cache miss, authenticate with another AcquireToken... method
		result, err = p.client.AcquireTokenByCredential(context.TODO(), scopes)
		if err != nil {
			return "", err
		}
	}

	return result.AccessToken, nil
}

func New(tenant, clientID, clientSecret string) (*Provider, error) {
	cred, err := confidential.NewCredFromSecret(clientSecret)
	if err != nil {
		panic(err)
	}
	client, err := confidential.New(
		"https://login.microsoftonline.com/"+tenant,
		clientID,
		cred,
	)
	if err != nil {
		return nil, err
	}

	provider := Provider{
		client: client,
	}

	_, err = provider.GetAccessToken()
	if err != nil {
		return nil, err
	}

	return &provider, nil
}

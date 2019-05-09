import {by, element} from 'protractor';

export class DexPage {
  getLoginInput(): any {
    return element(by.xpath('//input[@id="login"]'));
  }

  getPasswordInput(): any {
    return element(by.xpath('//input[@id="password"]'));
  }

  getLoginSubmitButton(): any {
    return element(by.id('submit-login'));
  }

  getLoginWithEmailButton(): any {
    return element(by.xpath('//span[contains(text(),"with Email")]'));
  }
}

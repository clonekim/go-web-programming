var PATTERNS =  {
  MONEY: /(^[+-]?\d+)(\d{3})/,
  PHONE: /^\d{10,11}$/,
  USERNAME: /^[a-zA-Z0-9]{4,20}$/,
  PASSWORD: /^.*(?=.{5,15})(?=.*\d)(?=.*[a-z|A-Z]).*$/,
  TELEPHONE: /(\d{3})(\d{3,4})(\d{4})/,
  ZIPCODE: /^[0-9]{5,6}$/,
  VERIFY_CODE: /^[0-9]{5}$/,
  NUMBER_2: /^[0-9]{2}$/,
  NUMBER_1MORE: /^[0-9]{1,}/,
  MONTH: /^[0-9]{2}$/,
  YEAR: /^[0-9]{4}$/,
  BIRTH: /^[0-9]{6}$/
}

function orderEventFormat(value) {
  switch(value) {
    case 'order':  return '신규/주문';
    case 'checkin': return '맡기기';
    case 'checkout': return '찾기';    
    default: return value;
  }
}

function dateShortFormat(value) {
  if(value)
    return dateFns.format(value, 'YYYY-MM-DD');
  else
    return '';
}

function dateLongFormat(value) {
  if(value)
    return dateFns.format(value, 'YYYY-MM-DD HH:mm');
  else
    return '';
}

function currency(value) {
  value += ''
  while(PATTERNS.MONEY.test(value)) {
    value = value.replace(PATTERNS.MONEY ,'$1' + ',' + '$2');
  }
  return value
};


function phoneFormat(value) {
  if (value)
    return value.replace(PATTERNS.TELEPHONE, '$1-$2-$3');
  else
    return value;
};

# credits: github.com/mebeim
# personalized a bit

import os
import sys
import re
from importlib import find_loader
from datetime import datetime, timedelta

def log(s, *a):
  sys.stderr.write('[advent] ' + s.format(*a))
  sys.stderr.flush()

def logcont(s, *a):
  sys.stderr.write(s.format(*a))
  sys.stderr.flush()

def check_or_die(resp):
  if resp.status_code != 200:
    logcont('\n')
    log('ERROR: response {}, url: {}\n', resp.status_code, resp.url)
    log('Did you log in and update your session cookie?\n')
    sys.exit(1)

  if 'please identify yourself' in resp.text.lower():
    logcont('\n')
    log('ERROR: Server returned 200, but is asking for identification.\n')
    log('Did you log in and update your session cookie?\n')
    sys.exit(1)

def check_setup_once():
  if YEAR == -1 and DAY == -1:
    now = datetime.utcnow() + timedelta(hours=-5)
    y, m, d = now.year, now.month, now.day

    if m != 12 or (m == 12 and d > 25):
      log('ERROR: year and day not set, and no event currently running!\n')
      sys.exit(1)

    log('Year and day not set, assuming today: Dec {}, {}.\n', d, y)
    setup(y, d)

def setup(year, day):
  global YEAR
  global DAY
  global SESSION

  if not (year >= 2015 and 1 <= day <= 25):
    log('ERROR: invalid year and/or day set!\n')
    sys.exit(1)

  YEAR = year
  DAY  = day

  if REQUESTS and os.path.isfile('secret_session_cookie'):
    with open('secret_session_cookie') as f:
      SESSION = f.read().rstrip()
      S.cookies.set('session', SESSION)

def get_input(fname=None, mode='r'):
  check_setup_once()

  if fname is not None:
    return open(fname, mode)

  if not os.path.isdir(CACHE_DIR.format(YEAR, DAY)):
    try:
      os.mkdir(CACHE_DIR.format(YEAR, DAY))
      log("Created cache directory '{}' since it did not exist.\n", CACHE_DIR.format(YEAR, DAY))
    except Exception as e:
      print(e)
      log("ERROR: could not create cache directory'{}'.\n", CACHE_DIR.format(YEAR, DAY))
      log('{}\n', str(e))
      sys.exit(1)

  log('Getting input for year {} day {}... ', YEAR, DAY)
  fname = os.path.join(CACHE_DIR.format(YEAR, DAY), 'input.txt')

  try:
    file = open(fname, mode)
    logcont('done (from disk).\n')
    return file
  except FileNotFoundError:
    pass

  if not REQUESTS:
    logcont('err!\n')
    log('ERROR: cannot download input, no requests module installed!\n')
    sys.exit(1)
  elif not SESSION:
    logcont('err!\n')
    log('ERROR: cannot download input file without session cookie!\n')
    sys.exit(1)

  logcont('downloading... ')

  r = S.get(URL.format(YEAR, DAY, 'input'))
  check_or_die(r)

  with open(fname, 'wb') as f:
    f.write(r.content)

  file = open(fname, mode)
  logcont('done.\n')

  return file

def print_answer(part, answer):
  print('Part {}:'.format(part), answer)

def wait(msg='Press [ENTER] to continue...'):
  '''Wait for user interaction by printing a message to standard error and
  waiting for input.
  '''
  log(msg)

  try:
    input()
  except KeyboardInterrupt:
    log(" keyboard interrupt, exiting...\n")
    sys.exit(0)

def submit_answer(part, answer):
  print_answer(part, answer)

  check_setup_once()

  if not REQUESTS:
    log('Cannot upload answer, no requests module installed!\n')
    print_answer(part, answer)
    return False

  wait('Press [ENTER] to submit, [CTRL+C] to cancel...')

  log('Submitting day {} part {} answer: {}\n', DAY, part, answer)

  r = S.post(URL.format(YEAR, DAY, 'answer'), data={'level': part, 'answer': answer})
  check_or_die(r)
  t = r.text.lower()

  if 'did you already complete it' in t:
    log('Already completed or wrong day/part.\n')
    return False

  if "that's the right answer" in t:
    matches = re.findall(r'rank\s+(\d+)', t)
    if matches:
      logcont('Right answer! Rank {}.\n', matches[0])
    else:
      log('Right answer!\n')

    if DAY == 25 and part == 1:
      log("It's Christmas! Automatically submitting second part...\n")
      S.post(URL.format(YEAR, 25, 'answer'), data={'level': 2, 'answer': 0})
      logcont('done!\n')
      log('Go check it out: https://adventofcode.com/{}/day/25#part2\n', YEAR)

    return True

  if 'you have to wait' in t:
    matches = re.compile(r'you have ([\w ]+) left to wait').findall(t)

    if matches:
      log('Submitting too fast, {} left to wait.\n', matches[0])
    else:
      log('Submitting too fast, slow down!\n')

    return False

  log('Wrong answer :(\n')
  return False

def get_lines(fin):
  lines = []
  for line in fin:
    lines.append(line.rstrip())
  return lines

def time(func):
  def wrapper():
    start_time = datetime.now()
    func()
    end_time = datetime.now()
    print('Duration: {}'.format(end_time - start_time))
  return wrapper

URL       = 'https://adventofcode.com/{:d}/day/{:d}/{:s}'
SESSION   = ''
CACHE_DIR = '../{:d}/day{:02d}/'
YEAR      = -1
DAY       = -1
REQUESTS  = find_loader('requests')

if REQUESTS:
  import requests
  S = requests.Session()
  S.headers['User-Agent'] = 'github.com/mebeim/aoc by marco AT mebeim.net'

# pyautogui
- 마우스 및 키보드를 자동으로 제어 할 수 있는 라이브러리

## Install
```bash
pip install pyautogui
```

## init
```python
import pyautogui

pag = pyautogui

print(pag.size()) # will shows your device monitor resolution
```


## mouse controlls
### moveTo
```pyton
# will moves your mouse pointer located in your monitor x=500, y=540
pag.moveTo(500, 540)
```
### click
```pyton
pag.click()
```

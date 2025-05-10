import cv2
import sys

screenshot_path = sys.argv[1]
template_path = sys.argv[2]

screenshot = cv2.imread(screenshot_path)
template = cv2.imread(template_path)
result = cv2.matchTemplate(screenshot, template, cv2.TM_CCOEFF_NORMED)

_, max_val, _, max_loc = cv2.minMaxLoc(result)

if max_val > 0.99:
    print("FOUND")
else:
    print("NOT FOUND")

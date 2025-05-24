# How the Two-Pointer Rain Water Trapping Algorithm Works

## Core Strategy

The algorithm uses **two pointers moving inward** and always processes the **shorter side first**. This guarantees we can safely calculate trapped water without knowing the complete picture.

## Why Process the Shorter Side?

**Key Insight:** Water level is limited by the shorter of two boundaries.

If `height[left] ≤ height[right]`, then:
- The right side is at least as tall as the left side
- Whatever happens on the right side later, it won't be shorter than current right
- So `leftMax` is the limiting factor for the left position
- We can safely calculate water at the left position

## Step-by-Step Process

Let's trace through `[3, 0, 2, 0, 4]`:

```
Initial: [3, 0, 2, 0, 4]
         L           R
leftMax = 0, rightMax = 0
```

### Step 1: Compare positions 0 and 4
```
height[0] = 3, height[4] = 4
3 ≤ 4, so process LEFT side

height[0] = 3 > leftMax = 0
→ Update leftMax = 3 (no water, just update max)
→ Move left pointer: left = 1
```

### Step 2: Compare positions 1 and 4  
```
height[1] = 0, height[4] = 4
0 ≤ 4, so process LEFT side

height[1] = 0 < leftMax = 3
→ Water trapped = 3 - 0 = 3 units
→ Move left pointer: left = 2
```

### Step 3: Compare positions 2 and 4
```
height[2] = 2, height[4] = 4  
2 ≤ 4, so process LEFT side

height[2] = 2 < leftMax = 3
→ Water trapped = 3 - 2 = 1 unit
→ Move left pointer: left = 3
```

### Step 4: Compare positions 3 and 4
```
height[3] = 0, height[4] = 4
0 ≤ 4, so process LEFT side

height[3] = 0 < leftMax = 3  
→ Water trapped = 3 - 0 = 3 units
→ Move left pointer: left = 4
```

### Step 5: Pointers meet
```
left = right = 4, loop ends
Total water = 3 + 1 + 3 = 7 units
```

## Visual Result
```
Original: [3, 0, 2, 0, 4]

|       |
|   |   |
|   |   |
|WWW|WWW|  ← 3+1+3 = 7 units of water
```

## Why This Works: The Guarantee

**The Magic:** When `height[left] ≤ height[right]`, we know:

1. **Current right boundary** is at least `height[left]` tall
2. **As we move right pointer inward**, it can only get taller or stay same
3. **So rightMax will always be ≥ height[left]**
4. **Therefore leftMax is the limiting factor** for current left position

This means we can calculate water at left position **without knowing future right values**!

## Algorithm Logic Flow

```
while left ≤ right:
    if height[left] ≤ height[right]:
        # Safe to process left side
        # rightMax will be at least height[left]
        
        if height[left] ≥ leftMax:
            leftMax = height[left]  # New peak, no water
        else:
            water += leftMax - height[left]  # Trap water
        
        left++
    else:
        # height[right] < height[left]
        # Safe to process right side
        # leftMax will be at least height[right]
        
        if height[right] ≥ rightMax:
            rightMax = height[right]  # New peak, no water  
        else:
            water += rightMax - height[right]  # Trap water
            
        right--
```

## Time & Space Complexity

- **Time:** O(n) - each element visited exactly once
- **Space:** O(1) - only using a few variables

Compare this to the brute force O(n²) approach where for each position you'd scan left and right to find the maximum heights!

## Key Takeaway

The brilliance is in the **guarantee**: by always processing the shorter side, we ensure that the "max" value we're using is indeed the limiting factor for that position, allowing us to make local decisions that build up to the global solution.
import '@testing-library/jest-dom'
import { expect } from 'vitest'
import * as matchers from 'vitest-axe/matchers'
import 'vitest-canvas-mock'

expect.extend(matchers)
